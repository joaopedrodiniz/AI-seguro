# 🤖 Assistente Virtual de Seguros de Vida com IA

Este projeto nasceu da vontade de unir tecnologia e experiência do dia a dia no setor de seguros. Atuando em uma seguradora de vida, percebi como a inteligência artificial pode ser uma aliada poderosa na hora de esclarecer dúvidas e melhorar o atendimento aos clientes.

Assim, surgiu a ideia de desenvolver um assistente virtual especializado em seguros de vida, utilizando IA local com o modelo Mistral via Ollama, sem dependência de serviços externos pagos.

O objetivo é criar uma interface acessível, capaz de responder com clareza e linguagem adequada às principais dúvidas sobre esse tipo de seguro.

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
