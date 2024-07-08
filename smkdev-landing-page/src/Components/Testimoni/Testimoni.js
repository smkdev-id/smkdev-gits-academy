import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';

const Testimoni = () => {
  return (
    <Container className='testimoni-section-wrapper p-5'>
      <Row className='testimoni-section d-grid justify-content-center align-items-center'>
        <Col className='testimoni-description'>
          <p className='text-center fs-2 fw-bold'>Apa Kata Mereka?</p>
          <p className='text-center fs-5'>
            Mereka telah merasakan serunya belajar skill digital di SMKDEV.
            Apakah kamu ingin seperti mereka?
          </p>
        </Col>
        <Col className='testimoni-card d-flex gap-3'>
          <Card className='card w-50 border-0 shadow-sm'>
            <Card.Body className='card-body'>
              <Card.Text className='fs-6 fst-italic text-muted'>
                SMKDEV ScholarshiCard.Text ini sangat menarik, karena memiliki
                penyampaian yang seru dan mudah dipahami. Selain itu juga di
                dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan
                mini games berhadiah
              </Card.Text>
              <Card.Title className='fs-5'>Rizal Maulana</Card.Title>
              <Card.Text className='card-subtitle mb-2 text-muted fst-italic'>
                SMK Negeri 1 Cimahi
              </Card.Text>
            </Card.Body>
          </Card>
          <Card className='card w-50 border-0 shadow-sm'>
            <Card.Body className='card-body'>
              <Card.Text className='fs-6 fst-italic text-muted'>
                Menjadi bagian keluarga SMKDev membantu mengembangkan
                keterampilan pemrograman saya (melalui event Coding Challenge
                misalnya). Selain itu, dukungan pemecahan masalah teknis,
                diskusi bersama, berbagi pengalaman (di Comunity Bounding
                misalnya) dan kesempatan berkolaborasi dalam sebuah proyek
                menjadi hal "menarik" yang tidak bisa saya lewatkan.
              </Card.Text>
              <Card.Title className='fs-5'>Asep Dwi Saputra</Card.Title>
              <Card.Text className='card-subtitle mb-2 text-muted fst-italic'>
                Teknik Informatika STMIK Widya Utama
              </Card.Text>
            </Card.Body>
          </Card>
          <Card className='card w-50 border-0 shadow-sm'>
            <Card.Body className='card-body'>
              <Card.Text className='fs-6 fst-italic text-muted'>
                Saya merasa sangat bahagia dan bangga karena berhasil menjadi
                pemenang dalam SMKDEV Coding Challenge. Selama kompetisi ini,
                menurut saya, pengalaman ini telah memperluas wawasan dan
                keterampilannya dalam pemrograman. Saya berterima kasih kepada
                tim penyelenggara SMKDEV karena memberikan kesempatan ini dan
                memberikan tantangan yang menginspirasi serta mengapresiasi
                kerja keras peserta lain.
              </Card.Text>
              <Card.Title className='fs-5'>Shevabey Rahman</Card.Title>
              <Card.Text className='card-subtitle mb-2 text-muted fst-italic'>
                Sistem Informasi Universitas Ahmad Dahlan
              </Card.Text>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Testimoni;
