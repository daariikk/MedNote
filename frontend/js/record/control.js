function pulseControl(indicator) {
    if (indicator < 50 || indicator > 100) return "CRITICAL";
    if ((indicator >= 50 && indicator <= 60) || (indicator >= 88 && indicator <= 100)) return "BAD";
    return "NORMAL";
  }


  function tensionsControl(upperPressure, lowerPressure) {
    if (upperPressure > 110 && upperPressure < 130 && lowerPressure > 75 && lowerPressure < 90) {
        return "NORMAL";
    }
    if (
        (upperPressure >= 100 && upperPressure <= 110 || upperPressure >= 130 && upperPressure <= 140) &&
        (lowerPressure >= 60 && lowerPressure <= 75 || lowerPressure >= 90 && lowerPressure <= 100)
    ) {
        return "BAD";
    }
    if (upperPressure < 100 || upperPressure > 140 || lowerPressure < 60 || lowerPressure > 100) {
        return "CRITICAL";
    }
    return "BAD"; 
}

function stepsControl(steps) {
    if (steps <= 3000) return "CRITICAL";
    if (steps > 3000 && steps < 10000) return "BAD";
    return "NORMAL";
}

function sleepControl(sleepHours, sleepMinutes) {
    const sum = sleepHours * 60 + sleepMinutes;
    if (sum <= 210) return "CRITICAL";
    if (sum > 210 && sum < 360) return "BAD";
    return "NORMAL";
}


function waterControl(waterVolume, glassesCount) {
    const volume = waterVolume * glassesCount; 
    if (volume < 800) return "CRITICAL";
    if (volume >= 800 && volume <= 1300) return "BAD";
    return "NORMAL";
}
