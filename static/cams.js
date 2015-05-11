function setScaleSize(scaleHeight) {
    var scales = document.getElementsByName("scale");
    for (var i = 0; i < scales.length; i++) {
        scales[i].style.height = scaleHeight;
    }
}

function setPosition(brandName, startPosition) {
    var cams = document.getElementsByName(brandName);
    for (var i = 0; i < cams.length; i++) {
        cams[i].style.top = startPosition + i * 20 + "px";
    }
    return cams.length * 20 + 20;
}

function enableBrand(brandName) {
    var cams = document.getElementsByName(brandName);
    for (var i = 0; i < cams.length; i++) {
        cams[i].style.visibility = "visible";
    }
}

function disableBrand(brandName) {
    var cams = document.getElementsByName(brandName);
    for (var i = 0; i < cams.length; i++) {
        cams[i].style.visibility = "hidden";
    }
}

function updateView() {
    var camSettings = document.getElementsByName("camVisible");
    var nextStartPosition = 30;
    for (var i = 0; i < camSettings.length; i++) {
        var brandName = camSettings[i].value;
        if (camSettings[i].checked == true) {
            nextStartPosition += setPosition(brandName, nextStartPosition);
            enableBrand(brandName);
        } else {
            disableBrand(brandName);
        }
    }
    setScaleSize(nextStartPosition - 20);
}

function toggleSettings() {
    var settings = document.getElementsByName("settings");
    if (settings[0].style.visibility == "visible") {
        disableBrand("settings");
    } else {
        enableBrand("settings");
    }
}