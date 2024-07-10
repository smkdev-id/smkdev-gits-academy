const MAX_POKEMON = 151;
const listWrapper = document.querySelector(".list-wrapper");
const searchInput = document.querySelector("#search-input");
const numberFilter = document.querySelector("#number");
const nameFilter = document.querySelector("#name");
const notFoundMessage = document.querySelector("#not-found-message");

//Mendeklarasikan sebuah array kosong untuk menyimpan data Pokemon
let allPokemons = [];

// Memanggil API PokeAPI untuk mendapatkan data Pokemon
fetch(`https://pokeapi.co/api/v2/pokemon?limit=${MAX_POKEMON}`)
  .then((response) => response.json())  // Mengubah response menjadi format JSON
  .then((data) => {
    allPokemons = data.results; // Menyimpan data Pokemon yang didapat dari API ke dalam array allPokemons
    // console.log(allPokemons)
    displayPokemons(allPokemons);
  });


// Mendeklarasikan fungsi asinkron untuk mengambil data Pokemon sebelum melakukan redirect
async function fetchPokemonDataBeforeRedirect(id) {
    try {
      // Menggunakan Promise.all untuk menjalankan dua fetch request secara paralel
      const [pokemon, pokemonSpecies] = await Promise.all([
        fetch(`https://pokeapi.co/api/v2/pokemon/${id}`).then((res) => res.json()),       // Mengambil data Pokemon berdasarkan ID dari API PokeAPI dan mengubahnya menjadi JSON
        fetch(`https://pokeapi.co/api/v2/pokemon-species/${id}`).then((res) => res.json()),     // Mengambil data spesies Pokemon berdasarkan ID dari API PokeAPI dan mengubahnya menjadi JSON
      ]);
      return true;    // Mengembalikan nilai true jika data berhasil diambil
    } catch (error) {
      // Menangkap dan menampilkan kesalahan jika terjadi masalah saat mengambil data
      console.error("Failed to fetch Pokemon data before redirect");
    }
  }
  

  // Fungsi untuk menampilkan daftar Pokemon dalam elemen listWrapper
function displayPokemons(pokemon) {
    // Mengosongkan konten dari listWrapper sebelum menambahkan elemen baru
    listWrapper.innerHTML = "";
  
    // Mengiterasi setiap objek Pokemon dalam array pokemon
    pokemon.forEach((pokemon) => {
      const pokemonID = pokemon.url.split("/")[6];   // Mendapatkan ID Pokemon dari URL dengan memisahkan string dan mengambil bagian ke-6
      const listItem = document.createElement("div"); // Membuat elemen div baru untuk setiap item dalam daftar
      listItem.className = "list-item"; // Menambahkan kelas CSS ke elemen div
      
      // Menetapkan konten HTML dalam elemen listItem
      listItem.innerHTML = `
          <div class="number-wrap">
              <p class="caption-fonts">#${pokemonID}</p>
          </div>
          <div class="img-wrap">
              <img src="https://raw.githubusercontent.com/pokeapi/sprites/master/sprites/pokemon/other/dream-world/${pokemonID}.svg" alt="${pokemon.name}" />
          </div>
          <div class="name-wrap">
              <p class="body3-fonts">#${pokemon.name}</p>
          </div>
      `;
  
      // Menambahkan event listener untuk menangani klik pada setiap item
      listItem.addEventListener("click", async () => {
        const success = await fetchPokemonDataBeforeRedirect(pokemonID);    // Memanggil fungsi fetchPokemonDataBeforeRedirect sebelum melakukan redirect
        if (success) {
          window.location.href = `./detail.html?id=${pokemonID}`;     // Jika data berhasil diambil, lakukan redirect ke halaman detail
        }
      });
      
      listWrapper.appendChild(listItem);  // Menambahkan elemen listItem ke dalam listWrapper
    });
  }


// Digunakan untuk menambahkan event listener pada elemen searchInput untuk event keyup,
// yang terjadi setiap kali pengguna melepaskan tombol keyboard setelah mengetik di input.
searchInput.addEventListener("keyup", handleSearch);


function handleSearch() {
    const searchTerm = searchInput.value.toLowerCase(); // Mendapatkan nilai input pencarian dan mengonversi ke huruf kecil
    let filteredPokemons; // Variabel untuk menyimpan hasil filter Pokemon
  
    // Memilih filter berdasarkan opsi yang dipilih (berdasarkan ID atau nama)
    if (numberFilter.checked) {
      filteredPokemons = allPokemons.filter((pokemon) => {
        const pokemonID = pokemon.url.split("/")[6]; // Mendapatkan ID Pokemon dari URL
        return pokemonID.startsWith(searchTerm); // Memeriksa apakah ID Pokemon dimulai dengan searchTerm
      });
    } else if (nameFilter.checked) {
      filteredPokemons = allPokemons.filter((pokemon) =>
        pokemon.name.toLowerCase().startsWith(searchTerm)
      );
    } else {
      filteredPokemons = allPokemons; // Jika tidak ada filter yang dipilih, tampilkan semua Pokemon
    }

    displayPokemons(filteredPokemons);  // Menampilkan Pokemon yang sudah difilter
  
    // Menampilkan atau menyembunyikan pesan jika Pokemon tidak ditemukan
    if (filteredPokemons.length === 0) {
      notFoundMessage.style.display = "block";
    } else {
      notFoundMessage.style.display = "none";
    }
  }


// Memilih elemen tombol close untuk pencarian
const closeButton = document.querySelector(".search-close-icon");

// Menambahkan event listener untuk klik pada tombol close
closeButton.addEventListener("click", clearSearch);

// Fungsi untuk membersihkan atau menghapus hasil pencarian
function clearSearch() {
  searchInput.value = ""; // Mengosongkan nilai input pencarian
  displayPokemons(allPokemons); // Menampilkan kembali semua Pokemon (tidak difilter)
  notFoundMessage.style.display = "none"; // Menyembunyikan pesan "not found" jika sebelumnya ditampilkan
}

  
  