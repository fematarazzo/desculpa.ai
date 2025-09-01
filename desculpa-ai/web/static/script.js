const themes = {
  dark: { body:'#1d2021', text:'#ebdbb2', navbar:'#3c3836', navText:'#ebdbb2', btn:'#fe8019' },
  light: { body:'#f2e5bc', text:'#3c3836', navbar:'#ebdbb2', navText:'#3c3836', btn:'#fe8019' },
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
