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
    document.getElementById('overlay').style.display = 'block';
    document.getElementById('popup').style.display = 'block';
}

function closePopup() {
    document.getElementById('overlay').style.display = 'none';
    document.getElementById('popup').style.display = 'none';
}
