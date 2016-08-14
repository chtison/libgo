class BaseConverterComponent {
    constructor() {
        this.CELL = `<input type="text" class="w3-input w3-border-light-grey w3-center" draggable="true" value="0123456789">`;
        this.number = document.getElementById('number');
        this.inBase = document.getElementById('inBase');
        this.toBase = document.getElementById('toBase');
        this.bases = document.getElementById('bases');
        this.btnAdd = document.getElementById('btnAdd');
        this.deleteArea = document.getElementById('deleteArea');
        const onDrop = (event) => { this.onDrop(event); };
        this.number.ondrop = onDrop;
        this.inBase.ondrop = onDrop;
        this.toBase.ondrop = onDrop;
        const deleteCell = (event) => { this.deleteCell(event); };
        this.deleteArea.ondrop = deleteCell;
        this.deleteArea.ondragover = () => { return false; };
        this.resetDeleteArea();
        this.btnAdd.addEventListener('click', () => { this.addCell('0123456789'); });
        this.initCells();
        this.bases.style.display = null;
    }
    initCells() {
        if (typeof Storage === 'undefined') {
            this.defaultCells();
            return;
        }
        let bases = window.localStorage.getItem('bases');
        if (bases === null || bases === "") {
            this.defaultCells();
            return;
        }
        bases = JSON.parse(bases);
        for (let i in bases) {
            this.addCell(bases[i]);
        }
    }
    defaultCells() {
        this.addCell('😃🌵🚀🍉🐳');
        this.addCell('0123456789');
        this.addCell('0123456789ABCDEF');
    }
    saveCells(event, index) {
        if (typeof Storage === 'undefined') {
            return;
        }
        if (index !== undefined) {
            --index;
        }
        let array = [];
        const inputs = this.bases.querySelectorAll('input');
        let value;
        for (let i = 0; i < inputs.length; ++i) {
            if (i === index)
                value = event.target.value;
            else
                value = inputs[i].getAttribute('value');
            array.push(value);
        }
        if (array.length === 0) {
            window.localStorage.removeItem('bases');
            return;
        }
        window.localStorage.setItem('bases', JSON.stringify(array));
    }
    resetDeleteArea() {
        this.deleteArea.innerText = 'Drag & Drop';
        this.deleteArea.classList.remove('w3-pale-red');
        const inputs = this.bases.querySelectorAll('input');
        for (let i = 0; i < inputs.length; ++i)
            inputs[i].removeAttribute('disabled');
    }
    onDragStart(event) {
        const inputs = this.bases.querySelectorAll('input');
        for (let i = 0; i < inputs.length; ++i)
            inputs[i].setAttribute('disabled', 'true');
        event.dataTransfer.setData('text/plain', event.target.value);
        event.dataTransfer.setData('index/plain', this.indexOfCell(event.target));
        this.deleteArea.innerText = 'Delete';
        this.deleteArea.classList.add('w3-pale-red');
    }
    onDrop(event) {
        event.preventDefault();
        event.target.value = event.dataTransfer.getData('text/plain');
    }
    deleteCell(event) {
        event.preventDefault();
        const i = Number(event.dataTransfer.getData('index/plain')) + 1;
        this.bases.querySelector(':nth-child(' + i + ')').remove();
        this.saveCells();
    }
    addCell(value) {
        const div = document.createElement('div');
        div.innerHTML = this.CELL;
        const cell = div.firstElementChild;
        cell.ondragstart = (event) => { this.onDragStart(event); };
        cell.ondragend = () => { this.resetDeleteArea(); };
        cell.setAttribute('value', value);
        cell.addEventListener('input', (event) => {
            this.saveCells(event, Number(this.indexOfCell(event.target)));
        });
        this.bases.insertBefore(div.firstChild, this.btnAdd);
        ;
        this.saveCells();
        return cell;
    }
    indexOfCell(cell) {
        const cells = this.bases.children;
        for (let i = 0; i < cells.length; ++i) {
            if (cell === cells[i]) {
                console.log('I:', i);
                return i.toString();
            }
        }
        console.log(cell, cells);
        return null;
    }
}
let baseConverterComponent;
window.onload = () => {
    baseConverterComponent = new BaseConverterComponent;
};
