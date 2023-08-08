const statusButton = document.getElementById('statusButton');
const timeDisplay = document.getElementById('timeDisplay');
let startTime = null;

statusButton.addEventListener('click', () => {
  if (statusButton.dataset.status === 'red') {
    statusButton.dataset.status = 'blue';
    statusButton.classList.remove('btn-primary');
    statusButton.classList.add('btn-info');
    statusButton.innerHTML = '<i class="fas fa-play"></i> Blau';
    startTime = new Date();
  } else {
    statusButton.dataset.status = 'red';
    statusButton.classList.remove('btn-info');
    statusButton.classList.add('btn-primary');
    statusButton.innerHTML = '<i class="fas fa-play"></i> Rot';
    const endTime = new Date();
    const timeDiff = endTime - startTime;
    const duration = formatDuration(timeDiff);
    const timestamp = formatTimestamp(startTime);
    const timeCard = createTimeCard(timestamp, duration);
    timeDisplay.appendChild(timeCard);
  }
});

function formatDuration(duration) {
  const milliseconds = duration % 1000;
  const seconds = Math.floor((duration / 1000) % 60);
  const minutes = Math.floor((duration / (1000 * 60)) % 60);
  const hours = Math.floor((duration / (1000 * 60 * 60)) % 24);
  return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}.${milliseconds.toString().padStart(3, '0')}`;
}

function formatTimestamp(timestamp) {
  const date = timestamp.toLocaleDateString('en-US');
  const time = timestamp.toLocaleTimeString('en-US');
  return `${date} ${time}`;
}

function createTimeCard(timestamp, duration) {
  const card = document.createElement('div');
  card.classList.add('card', 'mb-2');
  const cardBody = document.createElement('div');
  cardBody.classList.add('card-body');
  const cardTitle = document.createElement('h5');
  cardTitle.classList.add('card-title');
  cardTitle.textContent = timestamp;
  const cardText = document.createElement('p');
  cardText.classList.add('card-text');
  cardText.textContent = `Dauer: ${duration}`;
  cardBody.appendChild(cardTitle);
  cardBody.appendChild(cardText);
  card.appendChild(cardBody);
  return card;
}