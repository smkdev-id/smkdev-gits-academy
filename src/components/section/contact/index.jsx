import { IoIosBusiness, IoIosPin, IoMdCall } from "react-icons/io";
import Card from "./Card";
import Input from "../ui/Input";

const Contact = () => {
  const form = [
    { label: "Nama Anda", placeholder: "Type here ...", type: "text" },
    { label: "Alamat Email", placeholder: "Type here ...", type: "text" },
    { label: "Nomor WhatsApp", placeholder: "Type here ...", type: "text" },
    {
      label: "Nomor Satuan Pendidikan",
      placeholder: "Type here ...",
      type: "text",
    },
  ];

  const company = [
    {
      title: "Company information",
      value: "Creating High-Caliber Digital Talent",
      icon: IoIosBusiness,
    },
    {
      title: "Address",
      value:
        "Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295",
      icon: IoIosPin,
    },
    {
      title: "Contact us",
      value:
        "Email us for general queries, including marketing and partnership opportunities.",
      icon: IoMdCall,
    },
  ];

  return (
    <div className="py-10">
      <div className="flex flex-wrap w-full gap-5 md:flex-nowrap lg:flex-nowrap">
        {company.map((data, i) => {
          return (
            <Card
              key={i}
              title={data.title}
              value={data.value}
              Icon={data.icon}
            />
          );
        })}
      </div>
      <form className="flex flex-wrap w-full px-0 lg:px-32 md:px-32">
        {form.map((input, i) => (
          <div key={i} className="w-full p-3 md:w-1/2 lg:w-1/2">
            <Input
              label={input.label}
              placeholder={input.placeholder}
              type={input.type}
            />
          </div>
        ))}
        <label className="w-full form-control">
          <div className="px-4 label">
            <span className="text-xs font-semibold lg:text-sm md:text-sm label-text">Pesan</span>
          </div>
          <textarea
            className="m-3 textarea textarea-bordered textarea-lg focus:textarea-primary"
          ></textarea>
        </label>
      <button className="mx-3 text-white btn btn-primary">Send Message</button>
      </form>
    </div>
  );
};

export default Contact;
