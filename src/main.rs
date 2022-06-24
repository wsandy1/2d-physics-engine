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
    gravity: Vec2d<f64>,
    unit_size: f64,
    objects: Vec<PhysicsObject>,
}

impl PhysicsSolver {
    fn new(gravity: Vec2d<f64>, size: piston::window::Size) -> Self {
        let unit_size = size.width / NUM_UNITS as f64;
        return Self { gravity, unit_size, objects: vec![PhysicsObject { position_current: [5.0, 0.], position_old: [5.0, 0.0], acceleration: [0.0, 0.0] }]};
    }

    fn resize(&mut self, size: &piston::ResizeArgs) {
        let unit_size = size.window_size[0] / NUM_UNITS as f64;
        self.unit_size = unit_size;
    }

    fn update(&mut self, dt: f64) {
        // const SUB_STEPS: u32 = 16;
        // let sub_dt: f64 = dt / SUB_STEPS as f64;
        // for _ in 0..SUB_STEPS {
        //     self.apply_gravity();
        //     self.update_positions(sub_dt);
        // }
        self.apply_gravity();
        self.update_positions(dt);
    }

    fn apply_gravity(&mut self) {
        for obj in self.objects.iter_mut() {
            obj.accelerate(self.gravity);
        }
    }

    fn update_positions(&mut self, dt: f64) {
        for obj in self.objects.iter_mut() {
            obj.update_position(dt);
        }
    }
}

struct PhysicsObject {
    position_current: Vec2d<f64>,
    position_old: Vec2d<f64>,
    acceleration: Vec2d<f64>,
}

impl PhysicsObject {  

    fn get_pos(&self, unit_size: f64) -> Vec2d {
        return [self.position_current[0] * unit_size, self.position_current[1] * unit_size]
    }

    fn update_position(&mut self, dt: f64) {
        let velocity: Vec2d<f64> = sub(self.position_current, self.position_old);
        self.position_old = self.position_current;
        self.position_current = add(add(self.position_current, velocity), mul_scalar(self.acceleration, dt * dt));
        self.acceleration = [0.0, 0.0];
    }

    
    fn accelerate(&mut self, acc: Vec2d<f64>) {
        self.acceleration = add(self.acceleration, acc);
    }
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
            // polygon(
            //     WHITE,
            //     &[
            //         [ 0.0, 45.0],
            //         [ 0.0,  0.0],
            //         [15.0,  0.0],
            //         [15.0, 30.0],
            //         [30.0, 30.0],
            //         [30.0, 45.0],
            //     ],
            //     c.transform.trans(100f64, 100f64),
            //     gl,
            // );
            for obj in self.solver.objects.iter() {
                let trans = c.transform.trans(obj.get_pos(self.solver.unit_size)[0], obj.get_pos(self.solver.unit_size)[1]);
                rectangle(WHITE, rectangle::square(0.0, 0.0, 50.0), trans, gl);
            }
        });
    }

    fn resize(&mut self, args: &ResizeArgs) {
        self.solver.resize(args);
        println!("{}", self.solver.unit_size);
    }

    fn update(&mut self, args: &UpdateArgs) {
        self.solver.update(args.dt);
    }
}

fn main() {
    // Change this to OpenGL::V2_1 if not working.
    let opengl = OpenGL::V3_2;

    // Create a Glutin window.
    let mut window: GlutinWindow = WindowSettings::new("physics-sim", [1200, 800])
        .graphics_api(opengl)
        .exit_on_esc(true)
        .resizable(true)
        .build()
        .unwrap();

    let mut gi = GraphicsInterface {
        gl: GlGraphics::new(opengl),
        solver: PhysicsSolver::new([0.0, 9.8], window.size()),
    };

    let mut events = Events::new(EventSettings::new());
    while let Some(e) = events.next(&mut window) {
        if let Some(args) = e.render_args() {
            gi.render(&args);
        }

        if let Some(args) = e.update_args() {
            gi.update(&args);
        }

        if let Some(args) = e.resize_args() {
            gi.resize(&args);
        }

    }
}