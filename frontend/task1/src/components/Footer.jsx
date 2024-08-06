import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";

function FootComponent() {
  return (
    <main>
      <Container fluid className="bg-prim-2">
        <Container className=" py-5">
          <Navbar.Brand href="#home">
            <img src="https://smkdev.storage.googleapis.com/wp/SMKDEV-Logo-Long-150x38.png" height={35} alt="" />
            <h5 className="text-secondary mt-3 mb-4">Creating High-Caliber Digital Talent</h5>
            <p className="text-secondary fs-6">Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295</p>
          </Navbar.Brand>
        </Container>
      </Container>
      <Container className="text-center py-3">
        <h6 className="text-secondary">© 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.</h6>
      </Container>
    </main>
  );
}

export default FootComponent;
