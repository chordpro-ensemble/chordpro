use printpdf::*;
use std::fs;
use std::io::BufWriter;

mod lexer {
    pub enum Token {
        ChordL,
        ChordR,
        DirectiveL,
        DirectiveR,
    }
    pub fn process(song: String) -> String {
        return song;
    }
}

fn main() {
    let (doc, page1, layer1) = PdfDocument::new("Song Title", Mm(247.0), Mm(210.0), "Layer 1");
    let current_layer = doc.get_page(page1).get_layer(layer1);
    let font = doc.add_builtin_font(BuiltinFont::Helvetica).unwrap();

    let src = fs::read_to_string("examples/simple.cho").expect("error reading file");
    let processed = lexer::process(src);

    current_layer.use_text(processed, 11.0, Mm(10.0), Mm(200.0), &font);

    doc.save(&mut BufWriter::new(
        fs::File::create("test_working.pdf").unwrap(),
    ))
    .unwrap();
}
