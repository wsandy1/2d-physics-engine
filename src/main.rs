use glutin_window::GlutinWindow;
use opengl_graphics::{GlGraphics, OpenGL};
use piston::event_loop::{EventSettings, Events};
use piston::input::{RenderArgs, RenderEvent, UpdateArgs, UpdateEvent, ResizeEvent, ResizeArgs};
use piston::window::WindowSettings;
use piston::Window;
use graphics::math::*;


// here, 1 unit = 1 metre
const NUM_UNITS: u64 = 10;

struct PhysicsSolver {
    gravity: f64,
    unit_size: f64,
    objects: Vec<PhysicsObject>,
}

impl PhysicsSolver {
    fn new(gravity: f64, size: piston::window::Size) -> Self {
        let unit_size = size.width / NUM_UNITS as f64;
        return Self { gravity, unit_size, objects: vec![]};
    }

    fn resize(&mut self, size: &piston::ResizeArgs) {
        let unit_size = size.window_size[0] / NUM_UNITS as f64;
        self.unit_size = unit_size;
    }
}

struct PhysicsObject {
    position_current: Vec2d<f64>,
    position_old: Vec2d<f64>,
    acceleration: Vec2d<f64>,
}

pub struct GraphicsInterface {
    gl: GlGraphics, // OpenGL drawing backend.
    solver: PhysicsSolver,
}

impl GraphicsInterface {
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

    fn resize(&mut self, args: &ResizeArgs) {
        self.solver.resize(args);
        println!("{}", self.solver.unit_size);
    }
}

fn main() {
    // Change this to OpenGL::V2_1 if not working.
    let opengl = OpenGL::V3_2;

    // Create a Glutin window.
    let mut window: GlutinWindow = WindowSettings::new("spinning-square", [200, 200])
        .graphics_api(opengl)
        .exit_on_esc(true)
        .resizable(true)
        .build()
        .unwrap();

    let mut gi = GraphicsInterface {
        gl: GlGraphics::new(opengl),
        solver: PhysicsSolver::new(9.8, window.size()),
    };

    let mut events = Events::new(EventSettings::new());
    while let Some(e) = events.next(&mut window) {
        if let Some(args) = e.render_args() {
            gi.render(&args);
        }

        // if let Some(args) = e.update_args() {
        //     app.update(&args);
        // }

        if let Some(args) = e.resize_args() {
            gi.resize(&args);
        }

    }
}