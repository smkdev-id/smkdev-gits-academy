// Mendeklarasikan variabel global untuk menyimpan ID Pokemon saat ini
let currentPokemonId = null;

// Event listener untuk menangani saat DOM telah dimuat sepenuhnya
document.addEventListener("DOMContentLoaded", () => {
  const MAX_POKEMONS = 151;  // Konstanta untuk menyimpan jumlah maksimum Pokemon yang tersedia
  const pokemonID = new URLSearchParams(window.location.search).get("id");       // Mengambil ID Pokemon dari parameter query string pada URL
  const id = parseInt(pokemonID, 10);  // Mengonversi ID Pokemon dari string ke integer

  // Memeriksa apakah ID Pokemon valid
  if (id < 1 || id > MAX_POKEMONS) {
    return (window.location.href = "./index.html"); // Jika ID Pokemon tidak valid, arahkan kembali ke halaman index.html
  }
  currentPokemonId = id;   // Menyimpan ID Pokemon saat ini ke dalam variabel global currentPokemonId
  loadPokemon(id);  // Memuat data Pokemon dengan ID yang telah divalidasi
});

// Fungsi async untuk memuat data Pokemon berdasarkan ID
async function loadPokemon(id) {
    try {
      // Mengambil data Pokemon dan data spesies Pokemon secara bersamaan
      const [pokemon, pokemonSpecies] = await Promise.all([
        fetch(`https://pokeapi.co/api/v2/pokemon/${id}`).then((res) =>
          res.json()
        ),
        fetch(`https://pokeapi.co/api/v2/pokemon-species/${id}`).then((res) =>
          res.json()
        ),
      ]);
  
      const abilitiesWrapper = document.querySelector(  // Memilih wrapper untuk menampilkan kemampuan Pokemon
        ".pokemon-detail-wrap .pokemon-detail.move"
      );
      abilitiesWrapper.innerHTML = "";
  
      // Memeriksa apakah ID Pokemon yang dimuat saat ini sesuai dengan ID yang diminta
      if (currentPokemonId === id) {
        displayPokemonDetails(pokemon); // Menampilkan detail Pokemon
  
        const flavorText = getEnglishFlavorText(pokemonSpecies);    // Mendapatkan teks rasa (flavor text) dalam bahasa Inggris dari data spesies Pokemon
        document.querySelector(".body3-fonts.pokemon-description").textContent =
          flavorText;
  
        // Memilih panah navigasi kiri dan kanan
        const [leftArrow, rightArrow] = ["#leftArrow", "#rightArrow"].map((sel) =>
          document.querySelector(sel)
        );
  
        // Menghapus event listener yang ada sebelumnya untuk navigasi Pokemon
        leftArrow.removeEventListener("click", navigatePokemon);
        rightArrow.removeEventListener("click", navigatePokemon);
  
        // Menambahkan event listener untuk navigasi ke Pokemon sebelumnya jika tidak di halaman pertama
        if (id !== 1) {
          leftArrow.addEventListener("click", () => {
            navigatePokemon(id - 1);
          });
        }
  
        // Menambahkan event listener untuk navigasi ke Pokemon berikutnya jika tidak di halaman terakhir
        if (id !== 151) {
          rightArrow.addEventListener("click", () => {
            navigatePokemon(id + 1);
          });
        }
  
        window.history.pushState({}, "", `./detail.html?id=${id}`); // Mengubah URL pada history browser sesuai dengan ID Pokemon yang dimuat
      }
      return true;  // Mengembalikan nilai true jika proses pengambilan data berhasil
    } catch (error) {
      console.error("An error occured while fetching Pokemon data:", error);    // Menangani kesalahan jika gagal mengambil data Pokemon
      return false;     // Mengembalikan nilai false jika terjadi kesalahan
    }
  }

  // Definisi fungsi async untuk menavigasi ke Pokemon dengan ID baru
async function navigatePokemon(id) {
    currentPokemonId = id;  // Mengatur currentPokemonId ke ID yang baru
    await loadPokemon(id);      // Menunggu hingga proses pengambilan dan pemuatan data Pokemon selesai
  }


// Objek yang berisi warna background untuk setiap tipe Pokemon
const typeColors = {
    normal: "#A8A878",
    fire: "#F08030",
    water: "#6890F0",
    electric: "#F8D030",
    grass: "#78C850",
    ice: "#98D8D8",
    fighting: "#C03028",
    poison: "#A040A0",
    ground: "#E0C068",
    flying: "#A890F0",
    psychic: "#F85888",
    bug: "#A8B820",
    rock: "#B8A038",
    ghost: "#705898",
    dragon: "#7038F8",
    dark: "#705848",
    steel: "#B8B8D0",
  };
  
  // Fungsi untuk mengatur style elemen-elemen DOM
  function setElementStyles(elements, cssProperty, value) {
    elements.forEach((element) => {
      element.style[cssProperty] = value;
    });
  }
  
  // Fungsi untuk mengubah format hex color menjadi rgba
  function rgbaFromHex(hexColor) {
    return [
      parseInt(hexColor.slice(1, 3), 16),
      parseInt(hexColor.slice(3, 5), 16),
      parseInt(hexColor.slice(5, 7), 16),
    ].join(", ");
  }
  
  // Fungsi untuk mengatur warna background dan border berdasarkan tipe Pokemon
  function setTypeBackgroundColor(pokemon) {
    const mainType = pokemon.types[0].type.name; // Mendapatkan tipe Pokemon utama
    const color = typeColors[mainType]; // Mendapatkan warna berdasarkan tipe Pokemon
  
    if (!color) {
      console.warn(`Color not defined for type: ${mainType}`);
      return;
    }
  
    const detailMainElement = document.querySelector(".detail-main");
    setElementStyles([detailMainElement], "backgroundColor", color); // Mengatur background color
    setElementStyles([detailMainElement], "borderColor", color); // Mengatur border color
  
    // Mengatur background color untuk elemen-elemen tertentu
    setElementStyles(
      document.querySelectorAll(".power-wrapper > p"),
      "backgroundColor",
      color
    );
  
    // Mengatur color untuk teks pada elemen stats
    setElementStyles(
      document.querySelectorAll(".stats-wrap p.stats"),
      "color",
      color
    );
  
    // Mengatur color untuk progress bar pada elemen stats
    setElementStyles(
      document.querySelectorAll(".stats-wrap .progress-bar"),
      "color",
      color
    );
  
    // Mengubah warna background progress bar menggunakan CSS
    const rgbaColor = rgbaFromHex(color);
    const styleTag = document.createElement("style");
    styleTag.innerHTML = `
      .stats-wrap .progress-bar::-webkit-progress-bar {
          background-color: rgba(${rgbaColor}, 0.5);
      }
      .stats-wrap .progress-bar::-webkit-progress-value {
          background-color: ${color};
      }
    `;
    document.head.appendChild(styleTag); // Menambahkan style tag ke dalam head dokumen
  }
  
  // Fungsi untuk mengubah huruf pertama menjadi kapital
  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();
  }
  
  // Fungsi untuk membuat elemen baru dan menambahkannya sebagai child dari parent
  function createAndAppendElement(parent, tag, options = {}) {
    const element = document.createElement(tag);
    Object.keys(options).forEach((key) => {
      element[key] = options[key];
    });
    parent.appendChild(element);
    return element;
  }
  
  // Fungsi untuk menampilkan detail Pokemon pada halaman
  function displayPokemonDetails(pokemon) {
    const { name, id, types, weight, height, abilities, stats } = pokemon;
  
    // Mengubah judul halaman sesuai dengan nama Pokemon
    document.querySelector("title").textContent = capitalizeFirstLetter(name);
  
    const detailMainElement = document.querySelector(".detail-main");
    detailMainElement.classList.add(name.toLowerCase()); // Menambahkan class sesuai dengan nama Pokemon
  
    // Mengubah teks nama Pokemon pada halaman
    document.querySelector(".name-wrap .name").textContent =
      capitalizeFirstLetter(name);
  
    // Mengubah teks ID Pokemon pada halaman
    document.querySelector(".pokemon-id-wrap .body2-fonts").textContent = `#${String(
      id
    ).padStart(3, "0")}`;
  
    // Mengubah gambar Pokemon pada halaman
    const imageElement = document.querySelector(".detail-img-wrapper img");
    imageElement.src = `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/${id}.svg`;
    imageElement.alt = name;
  
    // Mengubah tampilan tipe Pokemon pada halaman
    const typeWrapper = document.querySelector(".power-wrapper");
    typeWrapper.innerHTML = "";
    types.forEach(({ type }) => {
      createAndAppendElement(typeWrapper, "p", {
        className: `body3-fonts type ${type.name}`,
        textContent: type.name,
      });
    });
  
    // Mengubah teks berat Pokemon pada halaman
    document.querySelector(
      ".pokemon-detail-wrap .pokemon-detail p.body3-fonts.weight"
    ).textContent = `${weight / 10}kg`;
  
    // Mengubah teks tinggi Pokemon pada halaman
    document.querySelector(
      ".pokemon-detail-wrap .pokemon-detail p.body3-fonts.height"
    ).textContent = `${height / 10}m`;
  
    // Mengubah daftar kemampuan Pokemon pada halaman
    const abilitiesWrapper = document.querySelector(
      ".pokemon-detail-wrap .pokemon-detail.move"
    );
    abilities.forEach(({ ability }) => {
      createAndAppendElement(abilitiesWrapper, "p", {
        className: "body3-fonts",
        textContent: ability.name,
      });
    });
  
    // Mengubah daftar stats Pokemon pada halaman
    const statsWrapper = document.querySelector(".stats-wrapper");
    statsWrapper.innerHTML = "";
  
    const statNameMapping = {
      hp: "HP",
      attack: "ATK",
      defense: "DEF",
      "special-attack": "SATK",
      "special-defense": "SDEF",
      speed: "SPD",
    };
  
    stats.forEach(({ stat, base_stat }) => {
      const statDiv = document.createElement("div");
      statDiv.className = "stats-wrap";
      statsWrapper.appendChild(statDiv);
  
      // Menambahkan teks nama stats Pokemon pada halaman
      createAndAppendElement(statDiv, "p", {
        className: "body3-fonts stats",
        textContent: statNameMapping[stat.name],
      });
  
      // Menambahkan teks base stat Pokemon pada halaman
      createAndAppendElement(statDiv, "p", {
        className: "body3-fonts",
        textContent: String(base_stat).padStart(3, "0"),
      });
  
      // Menambahkan progress bar stats Pokemon pada halaman
      createAndAppendElement(statDiv, "progress", {
        className: "progress-bar",
        value: base_stat,
        max: 100,
      });
    });
  
    // Mengatur warna background dan border berdasarkan tipe Pokemon
    setTypeBackgroundColor(pokemon);
  }
  
  // Fungsi untuk mendapatkan teks deskripsi Pokemon dalam bahasa Inggris
  function getEnglishFlavorText(pokemonSpecies) {
    for (let entry of pokemonSpecies.flavor_text_entries) {
      if (entry.language.name === "en") {
        let flavor = entry.flavor_text.replace(/\f/g, " ");
        return flavor;
      }
    }
    return "";
  }
  
  
  
