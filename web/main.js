document.addEventListener('DOMContentLoaded', function () {
  const lock = document.getElementById('tap-lock');
  lock.addEventListener('click', function () {
    fetch('/tap', {
      method: 'POST',
    }).then(res => console.log('Tap registered!'))
      .catch(() => console.error('Tap failed'));
  });
});
