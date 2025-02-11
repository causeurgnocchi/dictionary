document.querySelector('.search-box__submit').addEventListener('click', () => {
  const searchTerm = document.querySelector('.search-box__input').value;
  location.href = `/vocabularies/${searchTerm}`;
});