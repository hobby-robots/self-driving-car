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
