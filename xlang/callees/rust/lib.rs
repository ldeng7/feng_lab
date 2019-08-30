#[no_mangle]
pub extern fn rustFun(s: *mut u8, i: i32) -> i32 {
    unsafe {
        let mut v = Vec::from_raw_parts(s.offset(i as isize), 4, 4);
        v.copy_from_slice("rust".as_bytes());
        std::mem::forget(v);
    }
    i + 4
}
