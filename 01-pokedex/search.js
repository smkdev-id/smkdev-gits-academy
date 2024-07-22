// Seleksi elemen-elemen DOM yang diperlukan
const inputElement = document.querySelector("#search-input");
const search_icon = document.querySelector("#search-close-icon");
const sort_wrapper = document.querySelector(".sort-wrapper");

// Menambahkan event listener pada inputElement untuk meng-handle perubahan input
inputElement.addEventListener("input", () => {
  handleInputChange(inputElement);
});

// Menambahkan event listener pada search_icon untuk menangani klik pada ikon pencarian
search_icon.addEventListener("click", handleSearchCloseOnClick);

// Menambahkan event listener pada sort_wrapper untuk menangani klik pada ikon pengurutan
sort_wrapper.addEventListener("click", handleSortIconOnClick);

// Fungsi untuk meng-handle perubahan input
function handleInputChange(inputElement) {
  const inputValue = inputElement.value;

  // Menambahkan atau menghapus class untuk ikon pencarian berdasarkan nilai input
  if (inputValue !== "") {
    document
      .querySelector("#search-close-icon")
      .classList.add("search-close-icon-visible");
  } else {
    document
      .querySelector("#search-close-icon")
      .classList.remove("search-close-icon-visible");
  }
}

// Fungsi untuk meng-handle klik pada ikon penutup pencarian
function handleSearchCloseOnClick() {
  // Mengosongkan nilai input pada elemen dengan id "search-input"
  document.querySelector("#search-input").value = "";

  // Menghapus class "search-close-icon-visible" dari ikon pencarian
  document
    .querySelector("#search-close-icon")
    .classList.remove("search-close-icon-visible");
}

// Fungsi untuk meng-handle klik pada ikon pengurutan
function handleSortIconOnClick() {
  // Toggle class "filter-wrapper-open" pada elemen dengan class "filter-wrapper"
  document
    .querySelector(".filter-wrapper")
    .classList.toggle("filter-wrapper-open");

  // Toggle class "filter-wrapper-overlay" pada elemen <body>
  document.querySelector("body").classList.toggle("filter-wrapper-overlay");
}
