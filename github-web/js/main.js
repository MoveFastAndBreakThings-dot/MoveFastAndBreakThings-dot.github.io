(function () {
  const root = document.documentElement;
  const stored = localStorage.getItem('theme');
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  const theme = stored || (prefersDark ? 'dark' : 'light');
  root.setAttribute('data-theme', theme);

  const toggle = document.getElementById('theme-toggle');
  toggle.addEventListener('click', () => {
    const next = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
    root.setAttribute('data-theme', next);
    localStorage.setItem('theme', next);
  });

  const navToggle = document.getElementById('nav-toggle');
  const navLinks = document.getElementById('navbar-links');
  navToggle.addEventListener('click', () => navLinks.classList.toggle('open'));
  navLinks.querySelectorAll('a').forEach((a) =>
    a.addEventListener('click', () => navLinks.classList.remove('open'))
  );

  const progress = document.getElementById('progress');
  const setProgress = () => {
    const scrollTop = window.scrollY;
    const height = document.documentElement.scrollHeight - window.innerHeight;
    progress.value = height > 0 ? scrollTop / height : 0;
  };
  window.addEventListener('scroll', setProgress);
  window.addEventListener('resize', setProgress);
  setProgress();
})();
