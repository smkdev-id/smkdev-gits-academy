// navbar blur function
window.onscroll = function () {
  const header = document.querySelector("header");
  const fixedNav = header.offsetTop;

  if (window.scrollY > fixedNav) {
    header.classList.add("navbar-fixed");
  } else {
    header.classList.remove("navbar-fixed");
  }
};

// Hamburger
const hamburger = document.querySelector("#hamburger");
const navMenu = document.querySelector("#nav-menu");

hamburger.addEventListener("click", function () {
  hamburger.classList.toggle("hamburger-active");
  navMenu.classList.toggle("hidden");
});

// dynamic title
document.addEventListener("DOMContentLoaded", function () {
  const defaultTitle = "SMKDEV";
  const titles = {
    "#home": "SMKDEV",
    "#about": "SMKDEV - About",
    "#mitra": "SMKDEV - Mitra",
    "#blog": "SMKDEV - Blog",
    "#contact": "SMKDEV - Contact",
  };

  function updateTitle() {
    const hash = window.location.hash;
    document.title = titles[hash] || defaultTitle;
  }

  window.addEventListener("hashchange", updateTitle);
  updateTitle(); // otomatis update title jika hash berubah
});

// otomatis ganti title dan url ketika di scroll
document.addEventListener("DOMContentLoaded", function () {
  const sections = {
    home: {
      element: document.getElementById("home"),
      title: "SMKDEV",
      hash: "#home",
    },
    about: {
      element: document.getElementById("about"),
      title: "SMKDEV - About",
      hash: "#about",
    },
    mitra: {
      element: document.getElementById("mitra"),
      title: "SMKDEV - Mitra",
      hash: "#mitra",
    },
    blog: {
      element: document.getElementById("blog"),
      title: "SMKDEV - Blog",
      hash: "#blog",
    },
    contact: {
      element: document.getElementById("contact"),
      title: "SMKDEV - Contact",
      hash: "#contact",
    },
  };

  function updateTitleAndURL() {
    const scrollPosition = window.scrollY + window.innerHeight / 2;

    for (const sectionKey in sections) {
      const section = sections[sectionKey];
      const sectionTop = section.element.offsetTop;
      const sectionHeight = section.element.offsetHeight;

      if (
        scrollPosition >= sectionTop &&
        scrollPosition < sectionTop + sectionHeight
      ) {
        if (document.title !== section.title) {
          document.title = section.title;
          history.replaceState(null, "", section.hash);
        }
        break;
      }
    }
  }

  window.addEventListener("scroll", updateTitleAndURL);
  updateTitleAndURL();
});
