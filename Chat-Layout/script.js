// Função para rolar para o final do chat
function scrollToBottom() {
    const chatContainer = document.getElementById('chat-container');
    chatContainer.scrollTop = chatContainer.scrollHeight;
  }

  // Adicionar mensagem do usuário
  function addUserMessage(message) {
    const chatContainer = document.getElementById('chat-container');
    const messageDiv = document.createElement('div');
    messageDiv.className = 'message user';
    messageDiv.innerHTML = `
      <div class="message-content">${message}</div>
      <div class="avatar">
        <span class="user-icon">👤</span>
      </div>
    `;
    chatContainer.appendChild(messageDiv);
    scrollToBottom();
  }

  // Adicionar mensagem do bot
  function addBotMessage(message) {
    const chatContainer = document.getElementById('chat-container');
    const messageDiv = document.createElement('div');
    messageDiv.className = 'message bot';
    messageDiv.innerHTML = `
      <div class="avatar">
        <span class="bot-icon">🤖</span>
      </div>
      <div class="message-content">${message}</div>
    `;
    chatContainer.appendChild(messageDiv);
    scrollToBottom();
  }

  // Adicionar animação de carregamento
  function addLoadingIndicator() {
    const chatContainer = document.getElementById('chat-container');
    const loadingDiv = document.createElement('div');
    loadingDiv.className = 'message bot loading-message';
    loadingDiv.id = 'loading-indicator';
    loadingDiv.innerHTML = `
      <div class="avatar">
        <span class="bot-icon">🤖</span>
      </div>
      <div class="message-content">
        <div class="loading">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>
      </div>
    `;
    chatContainer.appendChild(loadingDiv);
    scrollToBottom();
  }

  // Remover animação de carregamento
  function removeLoadingIndicator() {
    const loadingIndicator = document.getElementById('loading-indicator');
    if (loadingIndicator) {
      loadingIndicator.remove();
    }
  }

  // Função principal para enviar pergunta
  async function perguntar() {
    const input = document.getElementById('pergunta');
    const pergunta = input.value.trim();
    
    if (!pergunta) return;
    
    // Limpar input
    input.value = '';
    
    // Adicionar mensagem do usuário
    addUserMessage(pergunta);
    
    // Mostrar indicador de carregamento
    addLoadingIndicator();
    
    try {
      const response = await fetch("http://localhost:3550/perguntar", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ question: pergunta })
      });
      
      const data = await response.json();
      
      // Remover indicador de carregamento
      removeLoadingIndicator();
      
      // Adicionar resposta do bot
      if (data.error) {
        addBotMessage(`Desculpe, ocorreu um erro: ${data.error}`);
      } else {
        addBotMessage(data.resposta);
      }
    } catch (error) {
      // Remover indicador de carregamento
      removeLoadingIndicator();
      
      // Mostrar mensagem de erro
      addBotMessage("Desculpe, não consegui processar sua pergunta. Por favor, tente novamente mais tarde.");
      console.error("Erro:", error);
    }
  }

  // Permitir envio com a tecla Enter
  function handleKeyPress(event) {
    if (event.key === 'Enter') {
      perguntar();
    }
  }

  // Rolar para o final quando a página carrega
  window.onload = scrollToBottom;