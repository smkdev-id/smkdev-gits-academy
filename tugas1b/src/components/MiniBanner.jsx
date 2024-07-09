import Container from "react-bootstrap/esm/Container";
import Button from 'react-bootstrap/Button';

const MiniBanner = () => {
    return (
        <div style={{ 
            backgroundImage: "url('https://thumbs4.imagebam.com/17/6c/68/MEUKXB9_t.jpg')",
            backgroundSize: 'cover',
            backgroundPosition: 'center',
            height: '30vh',
            width: '100%',
            position: 'relative',
            opacity: 0.5 }}>
            <Container>
                <div style={{ color:"#ffffff", alignItems:"center", display:"flex", height:"30vh"}} className="mx-3">
                    <h3>Temukan Potensi Penuh Anda</h3>
                    <p>
                        Bergabunglah dengan komunitas mahasiswa berbakat dan mitra industri kami untuk berkolaborasi dalam proyek TI dunia nyata. 
                        <div>
                            <Button variant="secondary" className="my-1">Pelajari lebih lanjut</Button>{' '}
                            <Button variant="outline-light" className="my-1">Daftar</Button>{' '}
                        </div>
                    </p>
                </div>
            </Container>
        </div>
    );
}

export default MiniBanner;
