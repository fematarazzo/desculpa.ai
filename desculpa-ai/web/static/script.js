const themes = {
  dark: { body:'#1d2021', text:'#ebdbb2', navbar:'#3c3836', navText:'#ebdbb2', btn:'#fe8019' },
  light: { body:'#f2e5bc', text:'#3c3836', navbar:'#ebdb2', navText:'#3c3836', btn:'#fe8019' },
  win95: { body:'#007f66', text:'#fff', navbar:'#000080', navText:'#fff', btn:'#c0c0c0' },
  xp: { body:'#d0e4f7', text:'#fff', navbar:'#3b5998', navText:'#fff', btn:'#aaccee' }
};

function applyTheme(name) {
  const t = themes[name];
  if (!t) return;

  document.body.style.backgroundColor = t.body;
  document.body.style.color = t.text;

  const spinner = document.getElementById('spinner');
  if (spinner) spinner.style.borderTopColor = t.btn;

  const nav = document.querySelector('.navbar');
  if (nav) {
    nav.style.backgroundColor = t.navbar;
    nav.querySelectorAll('a, button').forEach(el => el.style.color = t.navText);
  }

  document.querySelectorAll('.submit-btn').forEach(el => el.style.backgroundColor = t.btn);

  const responseBox = document.getElementById('response-box');
  const responseOutput = document.getElementById('response-output');
  if (responseBox && responseOutput) {
    responseBox.style.color = t.text;
    responseOutput.style.backgroundColor = t.navbar;
    responseOutput.style.color = t.text;
  }

  document.body.classList.toggle("win95-font", name === "win95");
  document.body.classList.toggle("xp-font", name === "xp");

  localStorage.setItem('theme', name);
}

const saved = localStorage.getItem('theme');
if (saved && themes[saved]) applyTheme(saved);

document.querySelectorAll('.theme-selector button').forEach(btn => {
  btn.addEventListener('click', () => applyTheme(btn.dataset.theme));
});

function setSpinnerVisible(visible) {
  const spinner = document.getElementById('spinner');
  if (!spinner) return;
  spinner.classList.toggle('hidden', !visible);
}

function appendToOutput(text) {
  const out = document.getElementById('response-output');
  if (!out) return;
  out.textContent += text;
}

function resetOutput() {
  const out = document.getElementById('response-output');
  if (!out) return;
  out.textContent = '';
}

function enableShareButton() {
  const btn = document.getElementById('share-whatsapp');
  if (!btn) return;
  btn.disabled = false;
}

function disableShareButton() {
  const btn = document.getElementById('share-whatsapp');
  if (!btn) return;
  btn.disabled = true;
}

document.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('prompt-form');
  if (!form) return;

  disableShareButton();

  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    resetOutput();
    setSpinnerVisible(true);
    disableShareButton();

    const promptEl = document.getElementById('prompt');
    const prompt = promptEl ? promptEl.value : '';

    try {
      const res = await fetch('/stream', {
        method: 'POST',
        body: new URLSearchParams({ prompt })
      });

      if (!res.ok) {
        appendToOutput(`Erro: ${res.status} ${res.statusText}`);
        setSpinnerVisible(false);
        return;
      }

      const reader = res.body.getReader();
      const decoder = new TextDecoder();

      while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        appendToOutput(decoder.decode(value, { stream: true }));
        window.scrollTo(0, document.body.scrollHeight);
      }

      setSpinnerVisible(false);
      enableShareButton();
    } catch (err) {
      appendToOutput('Erro ao conectar ao servidor: ' + err.message);
      setSpinnerVisible(false);
    }
  });

  const shareBtn = document.getElementById('share-whatsapp');
  if (shareBtn) {
    shareBtn.addEventListener('click', function () {
      const responseBox = document.getElementById('response-output');
      const text = responseBox ? responseBox.textContent.trim() : '';
      if (!text) {
        alert('Nenhum conteúdo disponível para compartilhar.');
        return;
      }
      const encodedText = encodeURIComponent(text);
      const whatsappUrl = `https://wa.me/?text=${encodedText}`;
      window.open(whatsappUrl, '_blank');
    });
  }
});
