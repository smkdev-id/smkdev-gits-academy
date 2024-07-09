import Container from "react-bootstrap/esm/Container";
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const Benefits = () => {
    return (
        <Container>
            <Row style={{ margin:"auto" }}>
                <Col>
                    <h2 style={{ margin:"auto", width:"40vw" }}>
                        Kembangkan Potensi Anda dengan Pembelajaran Berbasis Proyek
                    </h2>
                    <p style={{ margin:"auto", width:"40vw" }}>
                        Bergabunglah dengan komunitas kami dan mulailah perjalanan mengembangkan bakat Anda melalui pembelajaran berbasis proyek. Dengan akses ke berbagai studi kasus proyek TI dari industri, Anda akan memperoleh keterampilan dan pengalaman berharga di bawah bimbingan mentor.
                    </p>
                </Col>
            </Row>
        </Container>
    )
}

export default Benefits;