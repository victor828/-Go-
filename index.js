function disemvowel(str) {
  //Quitar Vocales
  let nstr = "";
  for (let i = 0; i < str.length; i++) {
    const e = str[i];
    if (
      e.toUpperCase() == "A" ||
      e.toUpperCase() == "E" ||
      e.toUpperCase() == "I" ||
      e.toUpperCase() == "O" ||
      e.toUpperCase() == "U"
    ) {
    } else {
      nstr += e;
    }
  }
  return nstr;
}

console.log(disemvowel("This website is for losers LOL!"));
