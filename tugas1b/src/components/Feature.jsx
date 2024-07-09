import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';
import Image from 'react-bootstrap/Image';

const Feature = () => {
    return (
        <Container>
            <Row style={{ justifyContent:"center", alignItems:"center" }}>
                <Col>
                    <p>
                        Temukan bagaimana siswa kami berkolaborasi dengan mitra industri untuk menyelesaikan proyek TI yang inovatif dengan sukses. Melalui pembelajaran langsung dan pengalaman dunia nyata, mereka mengembangkan keterampilan berharga dan mencapai hasil luar biasa. Bergabunglah bersama kami dalam merayakan pencapaian mereka dan terinspirasi oleh perjalanan mereka.
                    </p>
                </Col>
                <Col>
                    <h3>
                        Memberdayakan Siswa untuk Unggul dalam Proyek IT
                    </h3>
                    <Button variant="secondary" className="my-1">Pelajari lebih lanjut</Button>{' '}
                    <Button variant="outline-dark" className="my-1">Daftar</Button>{' '}
                </Col>
            </Row>
            <Row style={{ width:"40vw", justifyContent:"center", alignItems:"center", margin:"auto" }}>
                <Image src="https://thumbs4.imagebam.com/02/2b/a9/MEUKUQ9_t.png" alt="Professional-5-Steps-SMKDEV-Build-Digital-Talent-2.png" fluid />
            </Row>
        </Container>
    )
}

export default Feature;