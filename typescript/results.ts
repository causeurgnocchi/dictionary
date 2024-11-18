interface Vocabulary {
  furigana: Array<string>;
  writing: string;
  meanings: Array<string>;
}

function Vocabulary(element: Element) : Vocabulary {
  return {
    furigana: Array.from(element.querySelectorAll('.reading span')).map((characterReading) => characterReading.textContent || ''),
    writing: element.querySelector('.writing')?.textContent || '',
    meanings: Array.from(element.querySelectorAll('.meaning')).map((meaning) => meaning.textContent || ''),
  }
}

function createStagingElement(vocabulary: Vocabulary, chosenMeaning: string) : HTMLDivElement {
  const container = document.createElement('div');
  container.classList.add('staging-area');

  const readingWriting = document.createElement('div');
  readingWriting.classList.add('reading-writing');
  container.appendChild(readingWriting);

  const reading = document.createElement('p');
  reading.classList.add('reading');
  readingWriting.appendChild(reading);
  vocabulary.furigana.forEach((f) => {
    const characterReading = document.createElement('span');
    characterReading.classList.add(`reading-length-${f.length}`);
    characterReading.textContent = f;
    reading.appendChild(characterReading);
  });
  
  const writing = document.createElement('p');
  writing.classList.add('writing');
  writing.textContent = vocabulary.writing;
  readingWriting.appendChild(writing);

  const chosenMeaningElement = document.createElement('p');
  chosenMeaningElement.classList.add('meaning-chosen');
  chosenMeaningElement.textContent = chosenMeaning;
  container.appendChild(chosenMeaningElement);

  const meanings = document.createElement('ul');
  meanings.classList.add('meanings');
  container.appendChild(meanings);

  vocabulary.meanings.forEach((m) => {
    const meaning = document.createElement('li');
    meaning.classList.add('meaning');
    meaning.textContent = m;
    meanings.appendChild(meaning);
  });

  return container;
}


function createVocabularyElement(vocabulary: Vocabulary, isStagingArea: boolean) : HTMLDivElement {
  const container = document.createElement('div');
  container.classList.add(isStagingArea ? 'staging-area' : 'vocabulary');

  const readingWriting = document.createElement('div');
  readingWriting.classList.add('reading-writing');
  container.appendChild(readingWriting);

  const reading = document.createElement('p');
  reading.classList.add('reading');
  readingWriting.appendChild(reading);
  vocabulary.furigana.forEach((f) => {
    const characterReading = document.createElement('span');
    characterReading.classList.add(`reading-length-${f.length}`);
    characterReading.textContent = f;
    reading.appendChild(characterReading);
  });
  
  const writing = document.createElement('p');
  writing.classList.add('writing');
  writing.textContent = vocabulary.writing;
  readingWriting.appendChild(writing);

  const meanings = document.createElement('ul');
  meanings.classList.add('meanings');
  container.appendChild(meanings);

  vocabulary.meanings.forEach((m) => {
    const meaning = document.createElement('li');
    meaning.classList.add('meaning');
    meaning.textContent = m;
    meanings.appendChild(meaning);
  });

  return container;
}

let previousVocabulary = Vocabulary(document.querySelector('.staging-area'));

console.log(previousVocabulary);

document.querySelectorAll('.vocabulary').forEach((element) => {
  element.addEventListener('click', (e) => {
    const stagedElement = document.querySelector('.staging-area');
    if (stagedElement === null) return;
    
    const meaning = e.target as Element;
    if (!meaning.classList.contains('meaning')) return;
    
    const vocabulary = Vocabulary(element);
    
    stagedElement.replaceChildren(createVocabularyElement(vocabulary, true));
    element.replaceChildren(createVocabularyElement(previousVocabulary, false));

    previousVocabulary = vocabulary;
  })
});