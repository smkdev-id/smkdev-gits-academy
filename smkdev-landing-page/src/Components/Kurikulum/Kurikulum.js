import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Image from 'react-bootstrap/Image';
import Document from '../../Assets/Icons/Document.svg';
import Write from '../../Assets/Icons/Write.svg';
import Teach from '../../Assets/Icons/Teach.svg';

const Kurikulum = () => {
  return (
    <>
      <Container className='kurikulum-section-wrapper p-5'>
        <Col className='kurikulum-section d-flex justify-content-center align-items-center'>
          <Row className='kurikulum-card-container col-6 d-grid gap-3 p-3'>
            <Col>
              <Card className='card w-100 border-0 shadow-sm'>
                <Card.Body class='card-body d-flex'>
                  <Image src={Document} width={170} />
                  <div className='card-content d-grid p-3'>
                    <Card.Title class='card-title'>Kurikulum</Card.Title>
                    <Card.Text class='card-text text-secondary'>
                      Kurikulum berstandar industri global, senantiasa
                      diCard.Texterbaharui
                    </Card.Text>
                  </div>
                </Card.Body>
              </Card>
            </Col>
            <Col>
              <Card className='card w-100 border-0 shadow-sm'>
                <Card.Body class='card-body d-flex'>
                  <Image src={Write} width={150} />
                  <div className='card-content d-grid p-3'>
                    <Card.Title class='card-title'>Metode</Card.Title>
                    <Card.Text class='card-text text-secondary'>
                      Metode belajar project-based learning, dapat menjadi
                      portfolio siswa
                    </Card.Text>
                  </div>
                </Card.Body>
              </Card>
            </Col>
            <Col>
              <Card className='card w-100 border-0 shadow-sm'>
                <Card.Body class='card-body d-flex'>
                  <Image src={Teach} width={150} />
                  <div className='card-content d-grid p-3'>
                    <Card.Title class='card-title'>Mentor</Card.Title>
                    <Card.Text class='card-text text-secondary'>
                      Dibimbing langsung oleh mentor expert dari industri
                      digital
                    </Card.Text>
                  </div>
                </Card.Body>
              </Card>
            </Col>
          </Row>
          <Row className='kurikulum-content-container d-grid p-5'>
            <Col xs={12}>
              <p className='fw-bold fs-2'>
                MENGAPA HARUS MEMILIH BELAJAR DI{' '}
                <span className='text-primary'>SMKDEV</span>
              </p>
              <p className='fs-6 text-secondary'>
                Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan
                belajar teori dan capstone project yang dapat menjadi portofolio
                pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari
                expert mentor dari dunia industri.
              </p>
            </Col>
          </Row>
        </Col>
      </Container>
    </>
  );
};

export default Kurikulum;
