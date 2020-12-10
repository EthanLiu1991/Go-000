注意两点：
1.Wrap 的时候，需要 Wrap 统一的错误，如 code.NotFound，这样在上层可以直接用 errors.Is 来判断，且不需要和 Dao 层耦合
2.不要吞掉原始的错误信息，用 fmt.Sprintf 包含进来，这样打日志的时候方便判断到底是什么原因导致出了问题


dao: 

 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))


biz:

if errors.Is(err, code.NotFound} {

}