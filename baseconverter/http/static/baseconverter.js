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

        addCell();
        let cell = addCell();
        cell.attr('value', '0123456789ABCDEF');
        cell = addCell();
        cell.attr('value', '😃🌵🚀🍉🐳');

        bases.fadeIn();
});

function addCell() {
        const cell = $(CELL);
        cell[0].ondragstart = onDragStart;
        cell[0].ondragend = resetDeleteArea;
        btnAdd.before(cell);
        return cell;
}

function deleteCell(event) {
        event.preventDefault();
        const i = event.dataTransfer.getData('index/plain');
        $(bases).find(':eq('+i+')').remove();
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
