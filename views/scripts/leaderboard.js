function insertIndex() {
    const bodySection = document
        .getElementById('popup-leaderboard')
        .querySelectorAll('tbody')[0];
    for (let i = 0; i < bodySection.rows.length; i++) {
        const row = bodySection.rows[0];
        const newCell = row.insertCell(0);
        newCell.textContent = i + 1;
    }
}

function openPopup() {
    document.getElementById('overlay-leaderboard').style.display = 'block';
    document.getElementById('popup-leaderboard').style.display = 'block';
}

function closePopup() {
    document.getElementById('overlay-leaderboard').style.display = 'none';
    document.getElementById('popup-leaderboard').style.display = 'none';
}
