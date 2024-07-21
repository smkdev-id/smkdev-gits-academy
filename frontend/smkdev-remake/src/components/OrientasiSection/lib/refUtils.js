/**
 * Menambahkan elemen ke dalam referensi jika elemen tersebut belum ada di dalam referensi.
 *
 * @param {HTMLElement} el - Elemen yang akan ditambahkan ke dalam referensi.
 * @param {React.RefObject} refs - Referensi yang berisi array elemen-elemen.
 */
export const addToRefs = (el, refs) => {
  if (el && !refs.current.includes(el)) {
    refs.current.push(el);
  }
};
