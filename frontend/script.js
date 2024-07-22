window.addEventListener("scroll", function () {
  const header = document.getElementById("header");
  if (window.scrollY > 0) {
    header.classList.add("shadow-on-scroll");
  } else {
    header.classList.remove("shadow-on-scroll");
  }
});

// Daftar teks yang akan ditampilkan secara berurutan
const texts = ["Masa Depan Indonesia", "Global"];

// Element span yang akan diubah teksnya
const dynamicText = document.querySelector(".dynamic-text");

let index = 0;
let charIndex = 0;
let typingEffect = setInterval(type, 120); // Kecepatan penulisan per karakter

function type() {
  if (charIndex < texts[index].length) {
    dynamicText.textContent += texts[index].charAt(charIndex);
    charIndex++;
  } else {
    clearInterval(typingEffect);
    setTimeout(erase, 1000); // Tunggu sebelum menghapus
  }
}

function erase() {
  if (charIndex > 0) {
    dynamicText.textContent = texts[index].substring(0, charIndex - 1);
    charIndex--;
    setTimeout(erase, 50); // Kecepatan penghapusan per karakter
  } else {
    index = (index + 1) % texts.length; // Pindah ke teks berikutnya
    typingEffect = setInterval(type, 120); // Mulai menulis lagi
  }
}

const swiper = new Swiper(".swiper-container", {
  slidesPerView: 4, // Jumlah logo yang ditampilkan per view
  spaceBetween: 30, // Jarak antar slide
  loop: true, // Looping slide
  autoplay: {
    delay: 3000, // Delay antar slide dalam ms
    disableOnInteraction: false, // Autoplay tidak berhenti pada interaksi
  },
  breakpoints: {
    640: {
      slidesPerView: 2,
      spaceBetween: 20,
    },
    768: {
      slidesPerView: 3,
      spaceBetween: 30,
    },
    1024: {
      slidesPerView: 4,
      spaceBetween: 30,
    },
  },
});

//Hamburger and mobile menu
function MenuHandler(el, val) {
  let MainList = el.parentElement.parentElement.getElementsByTagName("ul")[0];
  let closeIcon =
    el.parentElement.parentElement.getElementsByClassName("close-m-menu")[0];
  let showIcon =
    el.parentElement.parentElement.getElementsByClassName("show-m-menu")[0];
  if (val) {
    MainList.classList.remove("hidden");
    el.classList.add("hidden");
    closeIcon.classList.remove("hidden");
  } else {
    showIcon.classList.remove("hidden");
    MainList.classList.add("hidden");
    el.classList.add("hidden");
  }
}
