function sendReq(val) {
  fetch('/' + val)
}

function left() {
  sendReq('left');
}

function forward() {
  sendReq('forward');
}

function right() {
  sendReq('right');
}

function backward() {
  sendReq('backward');
}

function resetSteering() {
  sendReq('resetSteering');
}

function resetThrust() {
  sendReq('resetThrust');
}

var LEFT = 37
var TOP = 38
var RIGHT = 40
var BOTTOM = 39

document.body.onkeydown = function(e) {
    console.warn();
    switch (e.keyCode) {
      case LEFT:
        left()
      case TOP:
        forward()
      case RIGHT:
        right()
      case BOTTOM:
        backward()
    }
};
