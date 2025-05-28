# 🤖 Assistente Virtual de Seguros de Vida com IA

Este é um projeto pessoal desenvolvido com o objetivo de explorar o uso de **inteligência artificial local** no contexto de **seguros de vida**, unindo aprendizado prático com uma aplicação real do mercado onde atuo atualmente.

Durante conversas na empresa onde trabalho (uma seguradora de vida), o tema de IA tem sido cada vez mais discutido. Isso me motivou a criar um assistente virtual que fosse capaz de responder, de forma clara e profissional, dúvidas sobre seguros de vida.

---

## 💡 Funcionalidades

- Chat interativo com IA local (sem custos de API)
- Respostas geradas com base no modelo **Mistral** usando o **Ollama**
- Interface web simples e intuitiva
- Foco no tema “seguros de vida” com linguagem adaptada

---

## 🛠️ Tecnologias Utilizadas

| Camada       | Tecnologia           |
| ------------ | -------------------- |
| Backend      | Go (Golang), Gin     |
| IA local     | Ollama + modelo Mistral |
| Frontend     | HTML, CSS, JavaScript |
| Versionamento| GitHub               |

---

## 🚀 Como rodar localmente

### 1. Instale o [Ollama](https://ollama.com/download)

Depois da instalação, no terminal, execute:

```bash
ollama pull mistral
```
### 2. Rode o Servidor GO:

```bash
go run main.go
```
