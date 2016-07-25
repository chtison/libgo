"use strict";

var number;
var inBase;
var toBase;

window.addEventListener('load', function() {

        number = document.getElementById('number');
        inBase = document.getElementById('inbase');
        toBase = document.getElementById('tobase');

        number.ondrop = onDrop;
        inBase.ondrop = onDrop;
        toBase.ondrop = onDrop;

        var draggables = document.querySelectorAll('input[draggable]');
        for (var i = 0; i < draggables.length; ++i) {
                draggables[i].ondragstart = onDrag;
        }
});

function onDrag(e) {
        e.dataTransfer.setData('text/plain', e.target.value);
}

function onDrop(e) {
        e.preventDefault();
        e.target.value = e.dataTransfer.getData('text');
}
