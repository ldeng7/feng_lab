#[link(name = "xlangc", kind = "dylib")]
extern { fn cFun(s: *mut u8, i: i32) -> i32; }
#[link(name = "xlanggo", kind = "dylib")]
extern { fn goFun(s: *mut u8, i: i32) -> i32; }
#[link(name = "xlangrust", kind = "dylib")]
extern { fn rustFun(s: *mut u8, i: i32) -> i32; }

static FUNCS: &[unsafe extern fn(*mut u8, i32) -> i32] = &[
    cFun,
    goFun,
    rustFun,
];

fn main() {
	for f in FUNCS {
        let mut s = String::from("rust calls xxxx");
        let l: i32;
        unsafe { l = f(s.as_mut_ptr(), 11); }
        println!("{}", s.get(.. (l as usize)).unwrap());
	}
}
