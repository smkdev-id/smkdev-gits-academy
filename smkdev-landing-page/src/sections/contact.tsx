import { ContactForm, MediaQuery } from "@/components";
import { NextPage } from "next";

interface Props {}

const Contact: NextPage<Props> = ({}) => {
  return (
    <div className="bg-bg2 flex w-full justify-center py-20">
      <MediaQuery className="flex flex-col items-center gap-x-10 lg:flex-row lg:items-start lg:justify-between">
        <div className="mb-20 flex w-full flex-col gap-y-5 lg:w-1/2">
          <h1 className="text-3xl font-semibold text-primary">Kontak SMKDEV</h1>
          <p>
            Berapa investasi untuk menggunakan SMKDEV? Anda bisa mulai dengan{" "}
            <span className="font-bold">GRATIS</span>.
          </p>
          <p>Silahkan isi form disebelah untuk informasi lebih lanjut.</p>
        </div>

        <div className="w-full lg:w-1/2">
          <ContactForm />
        </div>
      </MediaQuery>
    </div>
  );
};

export default Contact;
