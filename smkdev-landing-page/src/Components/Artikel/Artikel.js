import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Image from 'react-bootstrap/Image';
import Button from 'react-bootstrap/Button';
import Thumbnail1 from '../../Assets/Artikel/Thumbnail1.png';
import Thumbnail2 from '../../Assets/Artikel/Thumbnail2.png';

const Artikel = () => {
  return (
    <Container className='artikel-section-wrapper p-5'>
      <Row className='artikel-section d-grid justify-content-center align-items-center'>
        <Col>
          <Row className='artikel-header'>
            <p className='text-center fw-bold fs-3'>Artikel Pilihan SMKDEV</p>
          </Row>
        </Col>
        <Col className='justify-content-center align-items-center text-center'>
          <Row className='artikel-card d-flex'>
            <Col>
              <Card className='card w-100 border-0 shadow-sm text-start'>
                <Col>
                  <Image src={Thumbnail1} class='card-img-top' alt='...' />
                </Col>
                <Card.Body class='card-body'>
                  <Card.Title class='card-title fw-bold fs-5'>
                    Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus
                    AI
                  </Card.Title>
                  <Card.Text class='card-text'>
                    Kenapa prompt engineering penting untuk AI generatif seperti
                    ChatGPT dan Google Gemini? Pelajari keterampilan, prospe
                  </Card.Text>
                  <Card.Link className='text-decoration-none'>
                    Baca Lebih Lanjut
                  </Card.Link>
                </Card.Body>
              </Card>
            </Col>
            <Col>
              <Card className='card w-100 border-0 shadow-sm text-start'>
                <Card.Body class='card-body'>
                  <Card.Title class='card-title fw-bold fs-5'>
                    Panduan Pemula Menguasai AI Tools (Alat AI)
                  </Card.Title>
                  <Card.Text class='card-text'>
                    Pelajari cara menguasai alat AI dengan panduan pemula ini.
                    Temukan langkah-langkah, alat populer, dan tips untuk sukses
                    dalam memanfaatkan teknologi
                  </Card.Text>
                  <Card.Link className='text-decoration-none'>
                    Baca Lebih Lanjut
                  </Card.Link>
                </Card.Body>
              </Card>
            </Col>
            <Col>
              <Card className='card w-100 border-0 shadow-sm text-start'>
                <Col>
                  <Image src={Thumbnail2} class='card-img-top' alt='...' />
                </Col>
                <Card.Body class='card-body'>
                  <Card.Title class='card-title fw-bold fs-5'>
                    Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis
                  </Card.Title>
                  <Card.Text class='card-text'>
                    Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn
                    Rate? Churn rate adalah persentase pelanggan yang berhenti
                    menggunakan produk atau
                  </Card.Text>
                  <Card.Link className='text-decoration-none'>
                    Baca Lebih Lanjut
                  </Card.Link>
                </Card.Body>
              </Card>
            </Col>
            <Col>
              <Card className='card w-100 border-0 shadow-sm text-start'>
                <Card.Body class='card-body'>
                  <Card.Title class='card-title fw-bold fs-5'>
                    Mengoptimalkan ROI dengan Strategi Analisis Data yang
                    Efektif
                  </Card.Title>
                  <Card.Text class='card-text'>
                    Dalam era digital yang semakin maju, data telah menjadi
                    salah satu aset paling berharga bagi perusahaan. Namun,
                    nilai sejati
                  </Card.Text>
                  <Card.Link className='text-decoration-none'>
                    Baca Lebih Lanjut
                  </Card.Link>
                </Card.Body>
              </Card>
            </Col>
          </Row>
          <Button variant='primary' className='mt-3'>
            Lihat Semua Artikel
          </Button>
        </Col>
      </Row>
    </Container>
  );
};

export default Artikel;
