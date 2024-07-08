// import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import './Styles/Global.css';
import Navbar from './Components/Navbar/Navbar';
import Home from './Components/Home/Home';
import Mitra from './Components/Mitra/Mitra';
import Orientasi from './Components/Orientasi/Orientasi';
import Kurikulum from './Components/Kurikulum/Kurikulum';
import Talent from './Components/Talent/Talent';
import Testimoni from './Components/Testimoni/Testimoni';
import Mentor from './Components/Mentor/Mentor';
import Artikel from './Components/Artikel/Artikel';
import Footer from './Components/Footer/Footer';

function App() {
  return (
    <div className='App'>
      <Navbar />
      <Home />
      <Mitra />
      <Orientasi />
      <Kurikulum />
      <Talent />
      <Testimoni />
      <Mentor />
      {/* <Artikel /> */}
      {/* <Footer /> */}
    </div>
  );
}

<link
  rel='stylesheet'
  href='https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css'
  integrity='sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN'
  crossorigin='anonymous'
/>;

export default App;
