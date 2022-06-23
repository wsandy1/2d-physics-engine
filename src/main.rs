use glutin_window::GlutinWindow as Window;
use opengl_graphics::{GlGraphics, OpenGL};
use piston::event_loop::{EventSettings, Events};
use piston::input::{RenderArgs, RenderEvent, UpdateArgs, UpdateEvent};
use piston::window::WindowSettings;

pub struct App {
    gl: GlGraphics, // OpenGL drawing backend.
}

impl App {
    fn render(&mut self, args: &RenderArgs) {
        use graphics::{*, types::*};

        const BLACK: Color = [0.0, 0.0, 0.0, 1.0];
        const WHITE: Color = [1.0, 1.0, 1.0, 1.0];


        self.gl.draw(args.viewport(), |c, gl| {
            clear(BLACK, gl);
            polygon(
                WHITE,
                &[
                    [ 0.0, 45.0],
                    [ 0.0,  0.0],
                    [15.0,  0.0],
                    [15.0, 30.0],
                    [30.0, 30.0],
                    [30.0, 45.0],
                ],
                c.transform.trans(100f64, 100f64),
                gl,
            );
        });
    }

    fn update(&mut self, args: &UpdateArgs) {

    }
}

fn main() {
    // Change this to OpenGL::V2_1 if not working.
    let opengl = OpenGL::V3_2;

    // Create a Glutin window.
    let mut window: Window = WindowSettings::new("spinning-square", [200, 200])
        .graphics_api(opengl)
        .exit_on_esc(true)
        .resizable(true)
        .build()
        .unwrap();

    // Create a new game and run it.
    let mut app = App {
        gl: GlGraphics::new(opengl),
    };

    let mut events = Events::new(EventSettings::new());
    while let Some(e) = events.next(&mut window) {
        if let Some(args) = e.render_args() {
            app.render(&args);
        }

        if let Some(args) = e.update_args() {
            app.update(&args);
        }
    }
}