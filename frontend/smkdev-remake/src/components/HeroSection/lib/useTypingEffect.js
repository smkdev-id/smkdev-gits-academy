import { useState, useEffect } from "react";

/**
 * Hook untuk mengatur efek mengetik pada teks.
 *
 * @param {string[]} messages Array pesan yang akan ditampilkan dengan efek mengetik.
 * @param {number} initialSpeed Kecepatan awal mengetik dalam milidetik.
 * @returns {{ text: string, isTransitioning: boolean }} Objek yang berisi teks yang sedang diketik dan status transisi.
 */
const useTypingEffect = (messages, initialSpeed) => {
  /**
   * Teks yang sedang diketik.
   * @type {string}
   */
  const [text, setText] = useState("");

  /**
   * Indeks pesan saat ini dalam array messages.
   * @type {number}
   */
  const [currentMessageIndex, setCurrentMessageIndex] = useState(0);

  /**
   * Status apakah teks sedang dihapus.
   * @type {boolean}
   */
  const [isDeleting, setIsDeleting] = useState(false);

  /**
   * Kecepatan mengetik dalam milidetik.
   * @type {number}
   */
  const [typingSpeed, setTypingSpeed] = useState(initialSpeed);

  /**
   * Status apakah sedang dalam transisi antar pesan.
   * @type {boolean}
   */
  const [isTransitioning, setIsTransitioning] = useState(false);

  useEffect(() => {
    /**
     * Fungsi untuk mengatur logika mengetik dan menghapus teks.
     */
    const handleTyping = () => {
      const currentMessage = messages[currentMessageIndex];
      const targetText = isDeleting ? "" : currentMessage;

      let currentIndex = isDeleting ? text.length - 1 : text.length + 1;
      setText(targetText.substring(0, currentIndex));

      if (isDeleting && text === "") {
        setIsDeleting(false);
        setIsTransitioning(false);
        setCurrentMessageIndex(
          (prevIndex) => (prevIndex + 1) % messages.length
        );
      } else if (!isDeleting && text === currentMessage) {
        setTimeout(() => {
          setIsTransitioning(true);
          setTimeout(() => {
            setIsDeleting(true);
          }, 1000); // Delay sebelum mulai menghapus teks
        }, 2000);
      }
    };

    /**
     * Interval untuk menjalankan handleTyping sesuai dengan typingSpeed.
     * @type {number}
     */
    const interval = setInterval(handleTyping, typingSpeed);

    return () => clearInterval(interval);
  }, [text, isDeleting, currentMessageIndex, typingSpeed, messages]);

  useEffect(() => {
    setTypingSpeed(isDeleting ? 50 : initialSpeed); // Sesuaikan kecepatan saat menghapus
  }, [isDeleting, initialSpeed]);

  return { text, isTransitioning };
};

export default useTypingEffect;
