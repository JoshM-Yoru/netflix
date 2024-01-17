function scrollToTop() {
    window.scrollTo(0, 0);
}

function handleButtonClick(e) {
    const searchInput = document.getElementById('search')
    const clicked = e.target;
    let content = clicked.textContent;

    if (content === 'Catalog') {
        searchInput.setAttribute('hx-get', '/catalog?page=1');
    } else if (content === 'Watch List') {
        searchInput.setAttribute('hx-get', '/watch-list?page=1');
    }
    htmx.process(searchInput)

    if (clicked.tagName === 'BUTTON' || clicked.tagName === 'A') {
        if (searchInput) {
            searchInput.value = '';
        }
    }
}
document.addEventListener('click', handleButtonClick);
