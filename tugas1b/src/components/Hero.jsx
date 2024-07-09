import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';
import Button from 'react-bootstrap/Button';

const HeroApps = () => {
    return (
        <Container style={{ marginTop:"8vh", marginBottom:"8vh" }}>
            <Row style={{ justifyContent:"center", alignItems:"center" }}>
                <Col>
                    <h2>
                        Kembangkan Potensi Anda Melalui Pembelajaran Berbasis Proyek
                    </h2>
                    <p>
                        Bergabunglah dengan komunitas kami untuk mengembangkan keterampilan Anda dan berkolaborasi dalam proyek TI dunia nyata.
                    </p>
                    <Button variant="secondary" className="my-1">Pelajari lebih lanjut</Button>{' '}
                    <Button variant="outline-dark" className="my-1">Daftar</Button>{' '}
                </Col>
                <Col>
                    <Image style={{ width:"60vh" }} src="https://thumbs4.imagebam.com/65/f5/f8/MEUKU7P_t.jpg" alt="d00d654863fd42f8794e539864502a02.jpg" rounded/>
                </Col>
            </Row>
        </Container>
    )
}

export default HeroApps;