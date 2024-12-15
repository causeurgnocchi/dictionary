function Vocabulary(element) {
    return {
        furigana: Array.from(element.querySelectorAll('.furigana span')).map((characterReading) => characterReading.textContent || ''),
        writing: element.querySelector('.writing')?.textContent || '',
        meanings: Array.from(element.querySelectorAll('.meaning')).map((meaning) => meaning.textContent || ''),
    };
}
function createChosenVocabularyElement(vocabulary) {
    const element = document.createElement('div');
    element.classList.add('vocabulary', 'chosen-vocabulary');
    const furiganaWriting = document.createElement('div');
    furiganaWriting.classList.add('furigana-writing');
    element.appendChild(furiganaWriting);
    const furigana = document.createElement('p');
    furigana.classList.add('furigana');
    furiganaWriting.appendChild(furigana);
    vocabulary.furigana.forEach((f) => {
        const characterReading = document.createElement('span');
        characterReading.classList.add(`furigana-length-${f.length}`);
        characterReading.textContent = f;
        furigana.appendChild(characterReading);
    });
    const writing = document.createElement('p');
    writing.classList.add('writing');
    writing.textContent = vocabulary.writing;
    furiganaWriting.appendChild(writing);
    if (vocabulary.meanings.length > 0) {
        const chosenMeaningElement = document.createElement('p');
        chosenMeaningElement.classList.add('chosen-meaning');
        chosenMeaningElement.textContent = vocabulary.meanings[0];
        element.appendChild(chosenMeaningElement);
    }
    if (vocabulary.meanings.length > 1) {
        const meanings = document.createElement('ul');
        meanings.classList.add('meanings');
        element.appendChild(meanings);
        vocabulary.meanings.slice(1).forEach((m) => {
            const meaning = document.createElement('li');
            meaning.classList.add('meaning');
            meaning.textContent = m;
            meanings.appendChild(meaning);
        });
    }
    return element;
}
function createVocabularyElement(vocabulary) {
    const vocabularyElement = document.createElement('div');
    vocabularyElement.classList.add('vocabulary');
    const furiganaWriting = document.createElement('div');
    furiganaWriting.classList.add('furigana-writing');
    vocabularyElement.appendChild(furiganaWriting);
    const furigana = document.createElement('p');
    furigana.classList.add('reading');
    furiganaWriting.appendChild(furigana);
    vocabulary.furigana.forEach((f) => {
        const characterReading = document.createElement('span');
        characterReading.classList.add(`reading-length-${f.length}`);
        characterReading.textContent = f;
        furigana.appendChild(characterReading);
    });
    const writing = document.createElement('p');
    writing.classList.add('writing');
    writing.textContent = vocabulary.writing;
    furiganaWriting.appendChild(writing);
    if (vocabulary.meanings.length > 0) {
        const meanings = document.createElement('ul');
        meanings.classList.add('meanings');
        vocabularyElement.appendChild(meanings);
        vocabulary.meanings.forEach((m) => {
            const meaning = document.createElement('li');
            meaning.classList.add('meaning');
            meaning.textContent = m;
            meanings.appendChild(meaning);
        });
    }
    return vocabularyElement;
}
function getVocabulary(element) {
    return {
        writing: element.querySelector('.writing').textContent,
        furigana: [...element.querySelectorAll('.furigana span')].map((e) => e.textContent),
        meanings: [...element.querySelectorAll('.chosen-meaning, .meaning')].map((e) => e.textContent)
    };
}
document.querySelectorAll('.vocabulary').forEach((element) => {
    element.addEventListener('click', () => {
        console.log(getVocabulary(element));
    });
});
