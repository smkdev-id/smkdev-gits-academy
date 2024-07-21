import useTypingEffect from "./lib/useTypingEffect";

/**
 * Komponen untuk menampilkan teks dengan efek mengetik.
 *
 * @component
 * @returns {JSX.Element} Komponen TextTyping
 */
const TextTyping = () => {
  /**
   * Array pesan yang akan ditampilkan dengan efek mengetik.
   * @type {string[]}
   */
  const messages = ["Masa Depan Indonesia", "Global"];

  /**
   * Menggunakan hook useTypingEffect untuk mendapatkan teks yang sedang diketik dan status transisi.
   * @param {string[]} messages Pesan yang akan ditampilkan.
   * @param {number} speed Kecepatan mengetik dalam milidetik.
   * @returns {{ text: string, isTransitioning: boolean }} Objek yang berisi teks yang sedang diketik dan status transisi.
   */
  const { text, isTransitioning } = useTypingEffect(messages, 100);

  return (
    <div>
      <h1 className="text-5xl font-semibold text-custom-color-font font-poppins">
        Jadilah Talenta Digital{" "}
      </h1>
      <div className="relative inline-block">
        <p
          className={`text-5xl font-semibold text-custom-blue-secondary pt-5 font-poppins transition-opacity duration-500 ${
            isTransitioning ? "bg-custom-blue text-white" : ""
          }`}
        >
          {text}
        </p>
        {isTransitioning && <span className="absolute inset-0 "></span>}
      </div>
    </div>
  );
};

export default TextTyping;
