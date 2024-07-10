const wrapper = document.querySelector(".sliderWrapper");
const menuItems = document.querySelectorAll(".menu-item");
menuItems.forEach((item, index) => {
  item.addEventListener("click", () => {
    //Change the current slide
    wrapper.style.transform = `translateX(${-100 * index}vw)`;
  });
});
