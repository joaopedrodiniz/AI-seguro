package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Question string `json:"question"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaStreamResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// Função para carregar e montar o prompt com a pergunta
func loadPrompt(question string) (string, error) {
	data, err := ioutil.ReadFile("prompt.txt")
	if err != nil {
		return "", err
	}

	template := string(data)
	prompt := strings.Replace(template, "{{question}}", question, 1)
	return prompt, nil
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.Static("/", "./Chat-Layout")

	r.POST("/perguntar", func(c *gin.Context) {
		var body RequestBody

		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Println("Erro ao ler JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
			return
		}

		prompt, err := loadPrompt(body.Question)
		if err != nil {
			fmt.Println("Erro ao carregar prompt:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao preparar o prompt"})
			return
		}

		payload := OllamaRequest{
			Model:  "mistral",
			Prompt: prompt,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Erro ao criar payload JSON:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao preparar requisição"})
			return
		}

		resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Erro na requisição HTTP:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o modelo"})
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		var fullResponse string

		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Println("Linha de resposta do Ollama:", line)

			var streamResp OllamaStreamResponse
			if err := json.Unmarshal([]byte(line), &streamResp); err != nil {
				fmt.Println("Erro ao processar linha JSON:", err)
				continue
			}

			fullResponse += streamResp.Response
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Erro ao ler stream de resposta:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar resposta do modelo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"resposta": fullResponse})
	})

	fmt.Println("Servidor iniciado na porta 3550")
	if err := r.Run(":3550"); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
