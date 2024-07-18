function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

// Daftar kata-kata untuk ditekskan
const phrases = ["Global", "Masa Depan Indonesia"];
const el = document.getElementById("typewriter");

// Konfigurasi waktu delay
let sleepTime = 100;
let curPhraseIndex = 0;

// Fungsi untuk mengetikkan teks secara efek
const writeLoop = async () => {
  while (true) {
    let curWord = phrases[curPhraseIndex];

    // Mengetikkan kata secara bertahap
    for (let i = 0; i < curWord.length; i++) {
      el.innerText = curWord.substring(0, i + 1);
      await sleep(sleepTime);
    }

    await sleep(sleepTime * 10); // Delay setelah mengetikkan satu kata

    // Menghapus kata secara bertahap
    for (let i = curWord.length; i > 0; i--) {
      el.innerText = curWord.substring(0, i - 1);
      await sleep(sleepTime);
    }

    await sleep(sleepTime * 5); // Delay setelah menghapus satu kata

    // Memperbarui index untuk kata selanjutnya
    if (curPhraseIndex === phrases.length - 1) {
      curPhraseIndex = 0;
    } else {
      curPhraseIndex++;
    }
  }
};

// Menjalankan fungsi setelah dokumen dimuat sepenuhnya
document.addEventListener("DOMContentLoaded", writeLoop);
