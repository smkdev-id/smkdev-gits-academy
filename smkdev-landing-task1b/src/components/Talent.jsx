import TalentaImage from "../assets/miscellanious/talentadigital.png"

const Talent = () => {
    return (
        <>
            <div className="flex justify-center items-center px-5 py-12 bg-white">
                <div className="text-start md:text-center">
                    <h1 className="font-bold text-xl text-[#47bb8e]">Talenta Digital SMKDEV</h1>
                    <h1 className="font-bold md:text-3xl text-2xl mt-4">Mengasah Keterampilan Generasi Muda.<br/>Menyiapkan Talenta untuk Dunia Digital.</h1>
                    <img className="mt-4" src={TalentaImage} alt="TalentaImage" />
                </div>
                
            </div>
        </>
    )
}

export default Talent;