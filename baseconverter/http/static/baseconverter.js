"use strict";

const CELL = `<input type="text" class="w3-input w3-border-light-grey w3-center" draggable="true" value="0123456789">`;

var number;
var inBase;
var toBase;

var bases;
var btnAdd;
var deleteArea;

$(document).ready(function() {
        number = $('#number');
        inBase = $('#inbase');
        toBase = $('#tobase');

        number[0].ondrop = onDrop;
        inBase[0].ondrop = onDrop;
        toBase[0].ondrop = onDrop;

        bases = $('#bases');
        btnAdd = $('#btnAdd');
        deleteArea = $('#deletearea');

        deleteArea[0].ondrop = deleteCell;
        deleteArea[0].ondragover = onDragOver;
        resetDeleteArea();

        btnAdd.click(function(event) {
                addCell();
        });

        initCells();

        bases.fadeIn();
});

function initCells() {

        if (typeof(Storage) === "undefined") {
                defaultCells();
                return ;
        }

        let a = window.localStorage.getItem('bases');
        if (a === null || a === "") {
                defaultCells();
                return ;
        }
        a = JSON.parse(a);
        for (let i in a) {
                addCell(a[i]);
        }
}

function defaultCells() {
        addCell('0123456789');
        addCell('0123456789ABCDEF');
        addCell('😃🌵🚀🍉🐳');
}

function addCell(value) {
        const cell = $(CELL);
        cell[0].ondragstart = onDragStart;
        cell[0].ondragend = resetDeleteArea;
        cell.attr('value', value);
        btnAdd.before(cell);
        cell.on('input', function(event) { saveCells(event, $(this).index()); });
        saveCells();
        return cell;
}

function saveCells(event, index) {
        if (typeof(Storage) === "undefined") {
                return ;
        }
        if (index !== undefined) {
                --index;
        }
        let array = [];
        bases.children('input').map(function(i, elem) {
                let value = (i === index) ? event.target.value : elem.getAttribute('value');
                array.push(value);
        });
        if (array.length === 0) {
                window.localStorage.removeItem('bases');
                return ;
        }
        window.localStorage.setItem('bases', JSON.stringify(array));
}

function deleteCell(event) {
        event.preventDefault();
        const i = event.dataTransfer.getData('index/plain');
        $(bases).find(':eq('+i+')').remove();
        saveCells();
}

function onDragStart(event) {
        bases.find("input").attr('disabled', 'true');
        event.dataTransfer.setData('text/plain', event.target.value);
        event.dataTransfer.setData('index/plain', $(event.target).index());
        deleteArea.text('Delete');
        deleteArea.addClass('w3-pale-red');
}

function resetDeleteArea() {
        deleteArea.text('Drag & Drop');
        deleteArea.removeClass('w3-pale-red');
        bases.find("input").removeAttr('disabled');
}

function onDragOver(event) {
        if (event.target.id == 'deletearea') {
                event.preventDefault();
        }
}

function onDrop(event) {
        event.preventDefault();
        event.target.value = event.dataTransfer.getData('text');
}
