import { IoIosBusiness, IoIosPin, IoMdCall } from "react-icons/io";
import Card from "./Card";
import FormInput from "../ui/FormInput";

const Contact = () => {
  const company = [
    {
      title: "Company information",
      value:
        "Creating High-Caliber Digital Talent",
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
      <form>
        <FormInput label="tes" placeholder="tes" type="text" />
      </form>
    </div>
  );
};

export default Contact;
