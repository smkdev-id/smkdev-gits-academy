import React from "react";
import "./App.css";
// bootstrap
import NavComponent from "./components/Navbar";
import FootComponent from "./components/Footer";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Card from "react-bootstrap/Card";

// css
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  return (
    <div className="App">
      <NavComponent />
      <Container>
        <Row className="d-flex align-items-center py-5 my-5">
          <Col md="6">
            <h1>Jadilah Talenta Digital</h1>
            <h1 className="text-primary">Masa Depan Indonesia</h1>
            <h5 className="my-4 text-prim">Belajar langsung dengan expert dari industri dengan kurikulum komperhensif berbasis project-based learning</h5>
          </Col>
          <Col md="6">
            <img src="https://pisa.bantulkab.go.id/storage/pisa/child_information/DPOtUqoJ4IukHGNcLzB6KzLDwbsBeNoXHGdcbWxm.jpg" className="w-100 rounded-5" alt="ASD" />
          </Col>
        </Row>
      </Container>
      <Container fluid className="bg-prim-2 py-5">
        <Row className="text-center">
          <Col>
            <h4>Dipercaya Oleh Mitra Industri</h4>
            <img src="https://smkdev.storage.googleapis.com/wp/Logo-Gits-Indonesia.png" className="mx-2 " height={160} />
            <img src="	https://smkdev.storage.googleapis.com/wp/Logo-Eudeka.png" className="mx-2 " height={160} />
            <img src="https://smkdev.storage.googleapis.com/wp/logo-hitopia-4aaf6b-4acd41.png" className="mx-2 " height={160} />
            <img src="https://smkdev.storage.googleapis.com/wp/Logo-Arkana.png" className="mx-2 " height={160} />
            <img src="https://smkdev.storage.googleapis.com/wp/Logo-Mantab-One.png" className="mx-2 " height={160} />
          </Col>
        </Row>
      </Container>
      <Container className="py-5 my-5">
        <Row className=" d-flex justify-content-center">
          <Col md="8" className="text-center">
            <h2 className="text-prim">Orientasi Belajar SMKDEV</h2>
            <p className="my-4">Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri digital terkini.</p>
          </Col>
        </Row>
        <Row>
          <Col md="4">
            <Card className="rounded-4 shadow border-0">
              <Card.Body className="p-5 ">
                <h5 className="text-prim">Kurikulum Komprehensif</h5>
                <p>Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri</p>
              </Card.Body>
            </Card>
          </Col>
          <Col md="4">
            <Card className="rounded-4 shadow border-0">
              <Card.Body className="p-5 ">
                <h5 className="text-prim">Learning Path Industri</h5>
                <p>Learning path disusun berdasarkan kebutuhan dunia industri.</p>
              </Card.Body>
            </Card>
          </Col>

          <Col md="4">
            <Card className="rounded-4 shadow border-0 ">
              <Card.Body className="p-5 ">
                <h5 className="text-prim">Project-Based Learning</h5>
                <p>Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning</p>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
      <Container fluid className="bg-prim-2 py-5 my-5">
        <Container>
          <Row className="d-flex align-items-center">
            <Col md="6">
              <Card className="rounded-4 shadow border-0 my-4 ">
                <Card.Body className="p-5 ">
                  <h5 className="text-prim">Kurikulum</h5>
                  <p>Kurikulum berstandar industri global, senantiasa diperbaharui</p>
                </Card.Body>
              </Card>
              <Card className="rounded-4 shadow border-0 my-4 ">
                <Card.Body className="p-5 ">
                  <h5 className="text-prim">Metode</h5>
                  <p>Metode belajar project-based learning, dapat menjadi portfolio siswa</p>
                </Card.Body>
              </Card>
              <Card className="rounded-4 shadow border-0 my-4 ">
                <Card.Body className="p-5 ">
                  <h5 className="text-prim">Mentor</h5>
                  <p>Dibimbing langsung oleh mentor expert dari industri digital</p>
                </Card.Body>
              </Card>
            </Col>
            <Col md="6">
              <h2>
                MENGAPA HARUS MEMILIH BELAJAR DI <span className="text-primary">SMKDEV</span>
              </h2>
              <p>
                Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar teori dan capstone project yang dapat menjadi portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert mentor dari dunia industri.
              </p>
            </Col>
          </Row>
        </Container>
      </Container>
      <Container className="py-5 my-5">
        <Row className=" d-flex justify-content-center pb-5">
          <Col md="8" className="text-center">
            <h2 className="text-prim">Talenta Digital SMKDEV</h2>
            <img src="https://smkdev.storage.googleapis.com/wp/Professional-5-Steps-SMKDEV-Build-Digital-Talent-2.png" className="w-100" />
          </Col>
        </Row>
        <Row className=" d-flex justify-content-center">
          <Col md="8" className="text-center">
            <h2 className="text-prim">Apa Kata Mereka?</h2>
            <p className="my-4">Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah kamu ingin seperti mereka?</p>
          </Col>
        </Row>
        <Row>
          <Col md="4">
            <Card className="rounded-4 shadow border-0">
              <Card.Body className="p-5 ">
                <p>
                  Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai
                  peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst.
                </p>
                <h5 className="text-prim">Guruh Maulana</h5>
                <p>SMKN 11 Malang</p>
              </Card.Body>
            </Card>
          </Col>
          <Col md="4">
            <Card className="rounded-4 shadow border-0">
              <Card.Body className="p-5 ">
                <p>
                  Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah
                  banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!
                </p>
                <h5 className="text-prim">Tiara Febrianie</h5>
                <p>SMKN 11 Malang</p>
              </Card.Body>
            </Card>
          </Col>

          <Col md="4">
            <Card className="rounded-4 shadow border-0 ">
              <Card.Body className="p-5 ">
                <p>
                  Menjadi bagian keluarga SMKDev membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity
                  Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal "menarik" yang tidak bisa saya lewatkan.
                </p>
                <h5 className="text-prim">Asep Dwi Saputra</h5>
                <p>Teknik Inoformatika STMIK Widya Utama</p>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
      <FootComponent />
    </div>
  );
}

export default App;
