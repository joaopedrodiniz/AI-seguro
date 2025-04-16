package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

func main() {
	r := gin.Default()

	r.Static("/", "./Chat-Layout")

	r.POST("/perguntar", func(c *gin.Context) {
		var body RequestBody

		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Println("Erro ao ler JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
			return
		}

		prompt := fmt.Sprintf(`Você é um assistente especializado em seguros de vida. Sempre responda usando uma linguagem simples e profissional. Se a pergunta não estiver relacionada a seguros, diga educadamente que não pode ajudar.

Pergunta: %s`, body.Question)

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

		// Processar o stream de resposta linha por linha
		scanner := bufio.NewScanner(resp.Body)
		var fullResponse string

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Linha de resposta do Ollama:", line)

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

	fmt.Println("Servidor iniciado na porta 8080")
	if err := r.Run(":3550"); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
