var i = 0;
  var tulisan = "Jadilah Talenta Digital Masa Depan Indonesia";
  var kecepatan = 100; // Atur kecepatan mengetik (ms)
  var jedaAntaraKata = 1000; // Jeda sebelum mulai mengetik ulang (ms)

  function tulis() {
    // Pastikan id yang diambil sesuai dengan yang ada pada elemen HTML
    document.getElementById("lokasitulisan").textContent = tulisan.substring(0, i);

    if (i < tulisan.length) {
      i++;
      setTimeout(tulis, kecepatan);
    } else {
      setTimeout(resetTulis, jedaAntaraKata);
    }
  }

  function resetTulis() {
    i = 0;
    setTimeout(tulis, 0); // Mulai mengetik kembali setelah jeda
  }

  // Panggil fungsi tulis() untuk memulai efek mengetik
  window.onload = tulis;