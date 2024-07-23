/*=============== SHOW MENU ===============*/
const showMenu = (toggleId, navId) => {
  const toggle = document.getElementById(toggleId),
    nav = document.getElementById(navId);

  toggle.addEventListener('click', () => {
    nav.classList.toggle('show-menu');

    toggle.classList.toggle('show-icon');
  });
};

showMenu('nav-toggle', 'nav-menu');

/*=============== SCROLL ===============*/
const header = document.querySelector('header');
function toggleHeaderShadow() {
  if (window.scrollY > 0) {
    header.classList.add('shadow');
  } else {
    header.classList.remove('shadow');
  }
}

window.addEventListener('scroll', toggleHeaderShadow);

/*=============== TYPED ===============*/
let typed = new Typed('.auto__type', {
  strings: ['Global', 'Masa Depan Indonesia'],
  typeSpeed: 150,
  backSpeed: 150,
  loop: true,
});

/*=============== SWIPER ===============*/
const swiperMitra = new Swiper('.swiper__mitra', {
  direction: 'horizontal',
  loop: true,
  spaceBetween: 20,


  breakpoints: {
    360: {
      slidesPerView: 2,
    },
    640: {
      slidesPerView: 2,
    },
    768: {
      slidesPerView: 3,
    },
    1024: {
      slidesPerView: 5,
    },
  },
});

const swiperKata = new Swiper('.swiper__kata', {
  direction: 'horizontal',
  loop: true,
  spaceBetween: 20,


  breakpoints: {
    640: {
      slidesPerView: 2,
    },
    768: {
      slidesPerView: 2,
    },
    992: {
      slidesPerView: 2,
    },
    1024: {
      slidesPerView: 3,
    },
  },
});
