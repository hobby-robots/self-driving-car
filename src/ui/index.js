function sendReq(val) {
    fetch('/' + val)
}

function left() {
    if (!turning) {
        sendReq('left');
        turning = true
    }
}

function right() {
    if (!turning) {
        sendReq('right');
        turning = true
    }
}

function forward() {
    if (!moving) {
        sendReq('forward');
        moving = true
    }
}

function backward() {
    if (!moving){
        sendReq('backward');
        moving = true
    }
}

function resetSteering() {
    if (turning) sendReq('resetSteering');
    turning = false;
}

function resetThrust() {
    if (moving) sendReq('resetThrust');
    moving = false;
}

var turning = false;
var moving = false;

var LEFT = 37;
var FORWARD = 38;
var RIGHT = 39;
var BACKWARD = 40;

document.body.onkeydown = function (e) {
    switch (e.keyCode) {
        case LEFT:
            left()
            break
        case FORWARD:
            forward()
            break
        case RIGHT:
            right()
            break
        case BACKWARD:
            backward()
            break
    }
};

document.body.onkeyup = function (e) {
    switch (e.keyCode) {
        case LEFT:
            resetSteering()
            break
        case FORWARD:
            resetThrust()
            break
        case RIGHT:
            resetSteering()
            break
        case BACKWARD:
            resetThrust()
            break
    }
};