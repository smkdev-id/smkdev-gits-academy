/*Menu Toggle */
const menuToggle = document.querySelector('.toggle-menu input');
const nav = document.querySelector('nav ul');

menuToggle.addEventListener('click', function(){
    nav.classList.toggle('slide')
});