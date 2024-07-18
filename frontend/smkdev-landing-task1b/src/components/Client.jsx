import ArkanaLogo from "../assets/mitra/arkana.png"
import DigitsLogo from "../assets/mitra/digits.png"
import EudekaLogo from "../assets/mitra/eudeka.png"
import GitsLogo from "../assets/mitra/gits-indonesia.png"
import HitopiaLogo from "../assets/mitra/hitopia.png"
import MantabOneLogo from "../assets/mitra/mantab-one.png"

const Client = () => {
    return (
        <>
            <div className="text-start md:text-center bg-white px-5 py-12 mt-10">
                <h1 className="font-bold text-xl text-[#47bb8e]">Dipercaya Oleh Mitra Industri</h1>
                <h1 className="font-bold md:text-3xl text-2xl mt-4">Kualitas dan Keandalan Terjamin.<br/>
                Pilihan Utama Para Pemimpin Industri.</h1>
                {/* Desktop */}
                <div className="hidden md:grid grid-cols-3 justify-center px-4">
                    <div className="place-self-center">
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={ArkanaLogo} alt="ArkanaLogo"/>
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={DigitsLogo} alt="DigitsLogo"/>
                    </div>
                    <div className="place-self-center">
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={HitopiaLogo} alt="HitopiaLogo"/>
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={EudekaLogo} alt="EudekaLogo"/>
                    </div>
                    <div className="place-self-center">
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={GitsLogo} alt="GitsLogo"/>
                        <img className="md:h-[200px] w-[150px] md:w-auto" src={MantabOneLogo} alt="HitopiaLogo"/>
                    </div>
                </div>

                {/* Mobile View */}
                <div className="md:hidden grid grid-cols-2 justify-center px-4">
                    <div className="place-self-center">
                        <img className="w-[150px]" src={ArkanaLogo} alt="ArkanaLogo"/>
                        <img className="w-[150px]" src={DigitsLogo} alt="DigitsLogo"/>
                        <img className="w-[150px]" src={HitopiaLogo} alt="HitopiaLogo"/>
                    </div>
                    <div className="place-self-center">
                        <img className="w-[150px]" src={GitsLogo} alt="GitsLogo"/>
                        <img className="w-[150px]" src={MantabOneLogo} alt="HitopiaLogo"/>
                        <img className="w-[150px]" src={EudekaLogo} alt="EudekaLogo"/>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Client