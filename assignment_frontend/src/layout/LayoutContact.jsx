import React, { Fragment, useState } from "react";
import ContactColumn from "../components/contact/ContactColumn";
// import axios from "axios";
// import toast from "react-hot-toast";

const LayoutContact = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[90px] md:pb-[135px] px-[30px] md:px-[135px]";

  const [nama, setNama] = useState("");
  const [email, setEmail] = useState("");
  const [telepon, setTelepon] = useState("");
  const [instansi, setInstansi] = useState("");
  const [pesan, setPesan] = useState("");

  //   async function kirimPesan() {
  //     try {
  //       var gabungan =
  //         "Nama%3A%0A" +
  //         nama +
  //         "%0AEmail%3A%0A" +
  //         email +
  //         "%0APesan%3A%0A" +
  //         pesan;

  //       await axios
  //         .post(
  //           `https://api.telegram.org/bot${token}/sendMessage?chat_id=${grub}&text=${gabungan}&parse_mode=html`
  //         )
  //         .then((response) => {
  //           toast.success("Success Sending Message");
  //           console.log("Pesan berhasil dikirim!");
  //           // alert("berhasil")
  //           setNama("");
  //           setEmail("");
  //           setPesan("");
  //         });
  //     } catch (error) {
  //       console.error("Ada masalah dengan permintaan:", error);
  //     }
  //   }

  return (
    <Fragment>
      {/* <!-- CONTACT SECTION --> */}
      <section
        className={`${padding} w-full h-auto bg-gray-200 dark:bg-gray-900 font-poppins text-black dark:text-white`}
        id="contact"
      >
        <div className="container w-full h-full md:max-w-2xl mx-auto">
          {/* <!-- teks --> */}
          <div className="">
            <h4 className="text-center text-black dark:text-white font-bold text-3xl lg:text-5xl">
              Contact
            </h4>
            <h1 className="font-bold text-black dark:text-white text-center text-2xl pb-3">
              Contact Us
            </h1>
            <p className="text-center text-gray-600 dark:text-gray-400">
              Silahkan isi form di bawah untuk informasi lebih lanjut.
            </p>
          </div>
          {/* <!-- input --> */}
          <div className="w-full h-auto pt-20">
            <form>
              <ContactColumn
                title={"Nama"}
                type={"text"}
                value={nama}
                setWhat={setNama}
                placeholder={"Raturu"}
              />
              <ContactColumn
                title={"Email"}
                type={"email"}
                value={email}
                setWhat={setEmail}
                placeholder={"Raturu@gmail.com"}
              />
              <ContactColumn
                title={"Nomor Telepon"}
                type={"text"}
                value={telepon}
                setWhat={setTelepon}
                placeholder={"0881234567"}
              />
              <ContactColumn
                title={"Nama Instansi"}
                type={"text"}
                value={instansi}
                setWhat={setInstansi}
                placeholder={"Universtias Negeri Blababa"}
              />

              <div className="mb-3">
                <label
                  htmlFor="pesan"
                  className="font-bold text-black dark:text-white text-lg mb-3"
                >
                  Pesan
                </label>
                <br />
                <textarea
                  id="pesan"
                  value={pesan}
                  onChange={(e) => setPesan(e.target.value)}
                  cols="30"
                  rows="10"
                  className="text-black dark:text-white placeholder-gray-500 dark:placeholder-gray-400 w-full rounded-md border-2 border-solid border-gray-600 bg-transparent focus:outline-none focus:outline-primary focus:border-2 focus:ring-primary focus:outline-2 p-3"
                  placeholder="Tulis pesan Anda di sini..."
                ></textarea>
              </div>

              <button
                type="button"
                // onClick={kirimPesan}
                className="bg-blue-600 dark:bg-blue-800 transition-all duration-300 flex justify-center w-full py-2 rounded-full font-bold text-white hover:bg-blue-800 dark:hover:bg-blue-900 hover:cursor-pointer"
              >
                <h1>Kirim</h1>
              </button>
            </form>
          </div>
        </div>
      </section>
      {/* <!-- CONTACT SECTION END --> */}
    </Fragment>
  );
};

export default LayoutContact;
